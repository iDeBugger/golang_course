package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Profession string `json:"profession"`
}

var users map[int64]User = map[int64]User{}

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			processGet(w, req)
		case "POST":
			processPost(w, req)
		case "PUT":
			processPut(w, req)
		case "DELETE":
			processDelete(w, req)
		default:
			io.WriteString(w, "UNKNOWN METHOD")
		}
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func processGet(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.String()[1:], 10, 64) // url guaranteed to start from ASCII "/"
	if err != nil {
		io.WriteString(w, "Wrong user ID")
		return
	}
	user, ok := users[id]
	if !ok {
		io.WriteString(w, "User not found")
		return
	}
	fmt.Fprintf(w, "%v", user)
}

func processPost(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	user := User{}
	err := dec.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}
	users[user.ID] = user
	fmt.Fprintf(w, "Done")
}

func processPut(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.String()[1:], 10, 64) // url guaranteed to start from ASCII "/"
	if err != nil {
		io.WriteString(w, "Wrong user ID")
		return
	}
	dec := json.NewDecoder(req.Body)
	user := User{}
	err = dec.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}
	users[id] = user
	fmt.Fprintf(w, "Done")
}

func processDelete(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.String()[1:], 10, 64) // url guaranteed to start from ASCII "/"
	if err != nil {
		io.WriteString(w, "Wrong user ID")
		return
	}
	delete(users, id)
}
