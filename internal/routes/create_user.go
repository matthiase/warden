package routes

import (
	"encoding/json"
	"net/http"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	var data CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: validate the email address and name

	account, err := application.AccountStore.Create(data.Name, data.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: ensure the email address is not already registered

	// TODO: generate a random OTP

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateUserResponse{
		User: &User{
			Id:    account.Id,
			Email: account.Email,
			Name:  account.Name,
		},
		Otp: "123456",
	})
}
