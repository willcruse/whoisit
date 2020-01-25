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
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"pwd"`
}

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Hash string `json:"hash"`
}

type DatabaseInsert struct {
	UserEmail string `firestore:"email"`
	FirstName string `firestore:"firstName"`
	LastName  string `firestore:"lastName"`
	Hash      string `firestore:"hash"`
}

func NewUser(w http.ResponseWriter, r *http.Request) {

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
		resp := Response{false, err.Error(), ""}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			panic(err)
		}
		return
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "who-is-it-265713")
	if err != nil {
		resp := Response{false, err.Error(), ""}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			panic(err)
		}
		return
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(request.UserEmail).Get(ctx)
	if err == nil {
		resp := Response{true, "userExists", ""}
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	var databaseInsert DatabaseInsert
	databaseInsert.UserEmail = request.UserEmail
	databaseInsert.FirstName = request.FirstName
	databaseInsert.LastName = request.LastName
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		resp := Response{false, err.Error(), ""}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			panic(err)
		}
		return
	}

	databaseInsert.Hash = string(hash)
	_, err = client.Collection("users").Doc(databaseInsert.UserEmail).Set(ctx, databaseInsert)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		resp := Response{true, err.Error(), ""}
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			panic(err)
		}
		return
	}

	resp := Response{true, "", databaseInsert.Hash}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		resp := Response{false, err.Error(), ""}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			panic(err)
		}
	}
}
