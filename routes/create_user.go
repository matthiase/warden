package routes

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
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

	user, err := application.UserStore.Create(data.Name, data.Email)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	}

	passcode, err := application.PasscodeStore.Create(user.Id)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateUserResponse{
		User: &User{
			Id:    user.Id,
			Email: user.Email,
			Name:  user.Name,
		},
		Passcode: passcode.String(),
	})
}
