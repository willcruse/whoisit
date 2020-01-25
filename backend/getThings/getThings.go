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
	Things []Thing `json:"things"`
}

type Thing struct {
	User          string `firestore:"User"`
	Points        int    `firestore:"Score"`
	Thing         string `firestore:"Thing"`
	Justification string `firestore:"Justification"`
	ThingID       string `firestore:"ThingID"`
}

type Request struct {
	UserEmail string `json:"email"`
}

func GetThings(w http.ResponseWriter, r *http.Request) {
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
		resp := Response{make([]Thing, 0)}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			panic(err)
		}
		return
	}

	// Connect to firestore
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "who-is-it-265713")
	if err != nil {
		resp := Response{make([]Thing, 0)}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			log.Fatalln(err)
		}
	}
	defer client.Close()

	iter := client.Collection("things").Documents(ctx)
	resp := Response{make([]Thing, 0)}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			resp := Response{make([]Thing, 0)}
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(&resp); err != nil {
				log.Fatalln(err)
			}
		}
		var tempThing Thing
		if err := doc.DataTo(&tempThing); err != nil {
			resp := Response{make([]Thing, 0)}
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(&resp); err != nil {
				log.Fatalln(err)
			}
		}
		if tempThing.User == request.UserEmail {
			resp.Things = append(resp.Things, tempThing)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Fatalln(err)
	}
}
