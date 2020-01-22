package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

/*
Tweet model.
*/
// Tweet represents a tweet
type Tweet struct {
	ID       int       `json:"id"`
	UserName string    `json:"name"`
	Content  string    `json:"content"`
	Created  time.Time `json:"created"`
}

// Tweets represents a slice of Tweets.
type Tweets []Tweet

var tweetStore Tweets

func main() {

	r := newRouter()

	// Create http server and run.
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("listening on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

/*
REST Handlers
*/

// TweetList returns a list of tweet structs.
func tweetList(w http.ResponseWriter, r *http.Request) {
	tweets := tweetStore

	// Set Headers on the ResponseWriter
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Try to encode tweets as json and write to ResponseWriter. Panic if this returns an error.
	if err := json.NewEncoder(w).Encode(tweets); err != nil {
		panic(err)
	}
}

//TweetDetail returns a single tweet.
func tweetDetail(w http.ResponseWriter, r *http.Request) {
	var tweet Tweet
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	for _, t := range tweetStore {
		if t.ID == id {
			tweet = t
		}
	}

	// Set Headers on the ResponseWriter
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Try to encode tweets as json and write to ResponseWriter. Panic if this returns an error.
	if err := json.NewEncoder(w).Encode(tweet); err != nil {
		panic(err)
	}
}

// TweetCreate creates a new tweet and returns it.
func tweetCreate(w http.ResponseWriter, r *http.Request) {
	newTweet := Tweet{ID: len(tweetStore) + 1, UserName: "CodeWorkshop", Created: time.Now(), Content: "Big new news."}
	tweetStore = append(tweetStore, newTweet)
	// Set Headers on the ResponseWriter
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Try to encode tweets as json and write to ResponseWriter. Panic if this returns an error.
	if err := json.NewEncoder(w).Encode(newTweet); err != nil {
		panic(err)
	}
}

//TweetUpdate updates a tweet.
func tweetUpdate(w http.ResponseWriter, r *http.Request) {
	tweet := Tweet{UserName: "CodeWorkshop", Created: time.Now(), Content: "Some big polar bears here at the north pole today"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tweet); err != nil {
		panic(err)
	}
}

//TweetDelete updates a tweet.
func tweetDelete(w http.ResponseWriter, r *http.Request) {
	tweet := Tweet{UserName: "CodeWorkshop", Created: time.Now(), Content: "Some big polar bears here at the north pole today"}

	// Set Headers on the ResponseWriter
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Try to encode tweets as json and write to ResponseWriter. Panic if this returns an error.
	if err := json.NewEncoder(w).Encode(tweet); err != nil {
		panic(err)
	}
}

/*
Register handlers with the router.
*/

// NewRouter returns a mux router for our app.
func newRouter() *mux.Router {
	// Create a new mux router
	r := mux.NewRouter()
	//Register endpoints with the router to list the tweets.
	r.HandleFunc("/tweets", tweetList).
		Methods("GET")
	r.HandleFunc("/tweets/{id}", tweetDetail).
		Methods("GET")
	r.HandleFunc("/tweets", tweetCreate).
		Methods("POST")
	r.HandleFunc("/tweets/{id}", tweetUpdate).
		Methods("PUT")
	r.HandleFunc("/tweets/{id}", tweetDelete).
		Methods("DELETE")
	return r
}
