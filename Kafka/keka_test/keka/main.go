package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	mux.HandleFunc("GET /users/{id}", getUser)

	fmt.Println("Server listening to 8080")

	http.ListenAndServe(":8080", mux)
}


func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	
	//extract params received
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w, err.Error(), http.StatusBadRequest,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Swipe Successfully logged for id: ", id)
	w.WriteHeader(http.StatusOK)

}