package main

import (
	"fmt"
	"net/http"
)

var database = map[string]string{
	"user":  "user123",
	"password": "admin123",
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("user")
		password := r.FormValue("password")

		if pass, ok := database[user]; ok || pass == password{
			err := http.StatusUnauthorized
			http.Error(w, "Invalid username or password", err)
			return
		}

		next(w, r)
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("User %s Hit Endpoint", r.FormValue("user"))
		next(w, r)
	}
}

var middleware = []func(http.HandlerFunc) http.HandlerFunc{
	authMiddleware,
	loggingMiddleware,
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
}

func main() { 
	h := welcomeHandler
	for _, m := range middleware {
		h = m(h)
	}

	http.HandleFunc("/welcome", h)
	http.ListenAndServe(":8080", nil)
}