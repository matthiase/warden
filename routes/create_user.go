package routes

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateUserResponse struct {
	*User    `json:"user"`
	Passcode string `json:"passcode"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var data CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	}

	// TODO: validate the email address and name

	// TODO: ensure the email address is not already registered

	user, err := app.UserStore.Create(data.FirstName, data.LastName, data.Email)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	}

	passcode, err := app.PasscodeStore.Create(user.Id)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateUserResponse{
		User: &User{
			Id: user.Id,
		},
		Passcode: passcode,
	})
}
