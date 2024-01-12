package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/birdbox/authnz/identity"
	"github.com/birdbox/authnz/session"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateUserResponse struct {
	VerificationToken string `json:"verification_token"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var data CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	// TODO: validate the email address and name

	// TODO: ensure the email address is not already registered

	user, err := app.UserStore.Create(data.FirstName, data.LastName, data.Email)
	if err != nil {
		panic(err)
	}

	// Create a verification token for the user. This token will be used in
	// conjunction with the passcode to confirm the user's identity. Note that
	// the session claims are empty because the user has not yet been confirmed.
	identityClaims := identity.NewIdentityClaims(user.ID, &session.SessionClaims{}, app.Config)
	verificationToken, err := identityClaims.Sign([]byte(app.Config.Server.Secret))
	if err != nil {
		panic(err)
	}

	passcode, err := app.PasscodeStore.Create(user.ID)
	if err != nil {
		panic(err)
	}

	log.Printf("Created passcode %s for user %d", passcode, user.ID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateUserResponse{verificationToken})
}
