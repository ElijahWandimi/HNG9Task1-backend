package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", detailsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type details struct {
	slackUserName string
	backend       bool
	age           int
	bio           string
}

func loadPage() *details {
	return &details{slackUserName: "Elijah Wandimi", backend: true, age: 20, bio: "I like software engineering!"}
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	p := loadPage()
	fmt.Fprintf(w, "{slackUserName: %s, backend: %v, age: %d, bio: %s}", p.slackUserName, p.backend, p.age, p.bio)
}
