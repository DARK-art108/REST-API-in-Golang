package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{
		Name: "Ritesh",
		Age:  30,
	},
	{
		Name: "Raj",
		Age:  12,
	},
}

func main() {
	http.HandleFunc("/", serverstatus)
	http.HandleFunc("/users", userhandle)
	http.HandleFunc("/users/create", usercreatehandle)
	log.Printf("Starting server on port 8000...\n")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func serverstatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: %s\n", http.StatusText(http.StatusOK))
}

func userhandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func usercreatehandle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(users)
}
