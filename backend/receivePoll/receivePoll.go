package whoisit

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"
)

type Request struct {
	SubID  string `json:"subID"`
	Value  int    `json:"value"`
	UserID string `json:"userID"`
}

type Response struct {
	Success bool `json:"success"`
	Error string `json:"error"`
}

type Submission struct {
	User          string `firestore:"User"`
	Points        int    `firestore:"Score"`
	Thing         string `firestore:"Thing"`
	Justification string `firestore:"Justification"`
	PollScore     int    `firestore:"PollScore"`
	Votes         int    `firestore:"Votes"`
	SubID         string `firestore:"SubID"`
}

func ReceivePoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	if err := decoder.Decode(&request); err != nil {
		resp := Response{false, err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			panic(err)
		}
		return
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "who-is-it-265713")
	if err != nil {
		resp := Response{false, err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			panic(err)
		}
		return
	}
	defer client.Close()

	doc, err := client.Collection("submissions").Doc(request.SubID).Get(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "code = NotFound") {
			w.WriteHeader(http.StatusBadRequest)
			resp := Response{false, ""}
			if err := json.NewEncoder(w).Encode(&resp); err != nil {
				panic(err)
			}
			return
		}
		resp := Response{false, err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			panic(err)
		}
		return
	}

	var tempSub Submission
	if doc != nil && doc.Exists() {
		if err := doc.DataTo(&tempSub); err != nil {
			resp := Response{false, err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(&resp); err != nil {
				panic(err)
			}
			return
		}
	} else {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := Response{false, err.Error()}
			if err := json.NewEncoder(w).Encode(&resp); err != nil {
				panic(err)
			}
			return
		}
	}

	tempSub.PollScore += request.Value
	tempSub.Votes += 1

	if tempSub.Votes > 4 && tempSub.PollScore < 3 {
		_, err := client.Collection("submissions").Doc(request.SubID).Delete(ctx)
		if err != nil {
			resp := Response{false, err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(&resp); err != nil {
				panic(err)
			}
			return
		}
	} else {
		_, err = client.Collection("submissions").Doc(request.SubID).Set(ctx, tempSub)
		if err != nil {
			resp := Response{false, err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(&resp); err != nil {
				panic(err)
			}
			return
		}
	}
	resp := Response{true, ""}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		panic(err)
	}
}
