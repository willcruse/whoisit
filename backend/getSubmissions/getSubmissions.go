package whoisit

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Response struct {
	Submissions []Submission `json:"subs"`
}

type Submission struct {
	User          string `firestore:"User"`
	Points        int    `firestore:"Score"`
	Thing         string `firestore:"Thing"`
	Justification string `firestore:"Justification"`
	PollScore     int    `firestore:"PollScore"`
	Votes         int    `firestore:"Votes"`
	SubID string `firestore:"SubID"`
}

func GetSubmissions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Connect to firestore
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "who-is-it-265713")
	if err != nil {
		resp := Response{make([]Submission, 0)}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			log.Fatalln(err)
		}
	}
	defer client.Close()

	iter := client.Collection("submissions").Documents(ctx)
	resp := Response{make([]Submission, 0)}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			resp := Response{make([]Submission, 0)}
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(&resp); err != nil {
				log.Fatalln(err)
			}
		}
		var tempSub Submission
		if err := doc.DataTo(&tempSub); err != nil {
			resp := Response{make([]Submission, 0)}
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(&resp); err != nil {
				log.Fatalln(err)
			}
		}
		resp.Submissions = append(resp.Submissions, tempSub)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Fatalln(err)
	}
}
