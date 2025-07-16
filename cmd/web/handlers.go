package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("server", "go")
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// Add a snippetCreatePost handler function.
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// send 201 status code since something is created
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet..."))
}

// func snippetCategory(w http.ResponseWriter, r *http.Request) {
// 	category := r.PathValue("category")
// 	if category == "code" {
// 		w.Write([]byte("You are trying to view a code snippet."))
// 		return
// 	}
// 	w.Write([]byte("Category not found, specify a category of type 'code'."))
// }
