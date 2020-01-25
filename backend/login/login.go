package whoisit

import (
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
)

type Request struct {
	UserEmail string `json:"email"`
	Password  string `json:"pwd"`
}

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Login   bool   `json:"auth"`
	Hash    string `json:"hash"`
}

type User struct {
	UserEmail string `firestore:"email"`
	FirstName string `firestore:"firstName"`
	LastName  string `firestore:"lastName"`
	Hash      string `firestore:"hash"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	if err := decoder.Decode(&request); err != nil {
		resp := Response{false, err.Error(), false, ""}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			panic(err)
		}
		return
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "who-is-it-265713")
	if err != nil {
		resp := Response{false, err.Error(), false, ""}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			panic(err)
		}
		return
	}
	defer client.Close()

	doc, err := client.Collection("users").Doc(request.UserEmail).Get(ctx)
	if err != nil {
		resp := Response{true, "noUser", false, ""}
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	var user User
	if err = doc.DataTo(&user); err != nil {
		resp := Response{true, err.Error(), false, ""}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	var resp Response
	resp.Success = true
	resp.Error = ""
	if bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(request.Password)) == nil {
		resp.Login = true
		resp.Hash = user.Hash
	} else {
		resp.Login = false
		resp.Hash = ""
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		panic(err)
	}
}
