package routes

import (
	"encoding/json"
	"net/http"

	"github.com/birdbox/authnz/models"
)

type ConfirmUserRequest struct {
	UserId   int    `json:"user_id"`
	Passcode string `json:"passcode"`
}

type ConfirmUserResponse struct {
	*User        `json:"user"`
	AccessToken  string `json:"access_token"`
	SessionToken string `json:"session_token"`
}

func confirmUser(w http.ResponseWriter, r *http.Request) {
	var data ConfirmUserRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		BadRequestError(err.Error()).Render(w, r)
		return
	}

	// Look up the user id associated with the passcode
	passcode := models.Passcode(data.Passcode)
	userId, err := application.PasscodeStore.Find(passcode)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	}

	if userId != data.UserId {
		BadRequestError("Invalid passcode").Render(w, r)
		return
	}

	// Retrieve the user record
	user, err := application.UserStore.Find(userId)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	} else if user == nil {
		NotFoundError().Render(w, r)
		return
	}

	// Create the user session
	sessionToken, err := application.SessionStore.Create(user.Id)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ConfirmUserResponse{
		User: &User{
			Id:    user.Id,
			Email: user.Email,
			Name:  user.Name,
		},
		AccessToken:  "qwertyuiopasdfghjklzxcvbnm",
		SessionToken: sessionToken.String(),
	})
}
