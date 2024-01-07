package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/birdbox/authnz"
	"github.com/go-chi/chi/v5"
)

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

//type CreateUserResponse struct {
//	Email string `json:"email"`
//	Name  string `json:"name"`
//}

//var contextKey = authnz.ContextKey{Name: "user"}

func users(router chi.Router) {
	router.Post("/", createUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user authnz.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the email address and name
	fmt.Printf("%+v\n", user)
	w.WriteHeader(http.StatusCreated)
}

//func createUserContext(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		var user authnz.User
//		err := json.NewDecoder(r.Body).Decode(&user)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//
//		// Validate the email address and name
//
//		ctx := context.WithValue(r.Context(), contextKey, user)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
