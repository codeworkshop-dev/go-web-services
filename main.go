package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

func main() {

	// Setup the database and close when main thread exits.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := NewRouter()

	// Create http server and run.
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
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
func TweetList(w http.ResponseWriter, r *http.Request) {
	tweets := Tweets{
		Tweet{UserName: "CodeWorkshop", Created: time.Now(), Content: "Some big polar bears hear at the north pole today"},
		Tweet{UserName: "CodeWorkshop", Created: time.Now(), Content: "Saw over 5000 penguins, nothing new to report."},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Try to encode tweets as json and write to ResponseWriter. Panic if this returns an error.
	if err := json.NewEncoder(w).Encode(tweets); err != nil {
		panic(err)
	}
}

//TweetDetail returns a single tweet.
func TweetDetail(w http.ResponseWriter, r *http.Request) {
	tweet := Tweet{UserName: "CodeWorkshop", Created: time.Now(), Content: "Some big polar bears here at the north pole today"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tweet); err != nil {
		panic(err)
	}
}

// TweetCreate creates a new tweet and returns it.
func TweetCreate(w http.ResponseWriter, r *http.Request) {
	tweets := Tweets{
		Tweet{UserName: "CodeWorkshop", Created: time.Now(), Content: "Some big polar bears here at the north pole today"},
		Tweet{UserName: "CodeWorkshop", Created: time.Now(), Content: "Saw over 5000 penguins, nothing new to report."},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tweets); err != nil {
		panic(err)
	}
}

//TweetUpdate updates a tweet.
func TweetUpdate(w http.ResponseWriter, r *http.Request) {
	tweet := Tweet{UserName: "CodeWorkshop", Created: time.Now(), Content: "Some big polar bears here at the north pole today"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tweet); err != nil {
		panic(err)
	}
}

//TweetDelete updates a tweet.
func TweetDelete(w http.ResponseWriter, r *http.Request) {
	tweet := Tweet{UserName: "CodeWorkshop", Created: time.Now(), Content: "Some big polar bears here at the north pole today"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tweet); err != nil {
		panic(err)
	}
}

/*
Register handlers with the router.
*/

// NewRouter returns a mux router for our app.
func NewRouter() *mux.Router {
	// Create a new mux router
	r := mux.NewRouter()
	//Register an endpoint with the router to list the tweets.
	r.HandleFunc("/tweets", TweetList).
		Methods("GET")
	r.HandleFunc("/tweets/{id}", TweetDetail).
		Methods("GET")
	r.HandleFunc("/tweets", TweetList).
		Methods("PUT")
	r.HandleFunc("/tweets/{id}", TweetDetail).
		Methods("POST")
	r.HandleFunc("/tweets/{id}", TweetDelete).
		Methods("DELETE")
	return r
}

/*
Tweet model.
*/
// Tweet represents a tweet
type Tweet struct {
	UserName string    `json:"name"`
	Content  string    `json:"content"`
	Created  time.Time `json:"created"`
}

// Tweets represents a list of Tweets.
type Tweets []Tweet
