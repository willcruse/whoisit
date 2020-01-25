package whoisit

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"context"

	"cloud.google.com/go/firestore"
)

type Request struct {
	User          string `json:"userID"`
	Thing         string `json:"thing"`
	Score         int    `json:"score"`
	Justification string `json:"just"`
	PollScore int `firestore:"PollScore"` 
	Votes int `firestore:"Votes"`
	CurrSubID string `firestore:"SubID"`
}

type Response struct {
	Success bool `json:"success"`
}

func SubmitScore(w http.ResponseWriter, r *http.Request) {

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	request := new(Request)

	if err := decoder.Decode(&request); err != nil {
		log.Fatalln(err)
	}

	request.PollScore = 0
	request.Votes = 0

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "who-is-it-265713")
	if err != nil {
		resp := Response{true}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Fatalln(err)
	}
	}
	defer client.Close()

	//get next subID
	type subID struct {
		SubID int `firestore:"id"`
	}

	var nextSubID subID
	subIDDoc, err := client.Collection("submissions").Doc("nextSubID").Get(ctx)

	if err != nil {
		resp := Response{false}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			log.Fatalln(err)
		}
	}

	if err := subIDDoc.DataTo(&nextSubID); err != nil {
		resp := Response{false}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			log.Fatalln(err)
		}
	}

	toSubID := fmt.Sprintf("sub%d", nextSubID.SubID)
	request.CurrSubID = toSubID

	nextSubID.SubID = nextSubID.SubID + 1

	//write submission
	_, err = client.Collection("submissions").Doc(toSubID).Set(ctx, request)

	if err != nil {
		resp := Response{false}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			log.Fatalln(err)
		}
	}

	_, err = client.Collection("submissions").Doc("nextSubID").Set(ctx, nextSubID)

	//write success
	resp := Response{true}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		resp := Response{false}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			log.Fatalln(err)
		}
	}
}
