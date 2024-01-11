package routes

import (
	"encoding/json"
	"net/http"

	"github.com/birdbox/authnz/identity"
	"github.com/birdbox/authnz/session"
)

type ConfirmUserRequest struct {
	UserID   int    `json:"user_id"`
	Passcode string `json:"passcode"`
}

type ConfirmUserResponse struct {
	*User       `json:"user"`
	AccessToken string `json:"access_token"`
}

func confirmUser(w http.ResponseWriter, r *http.Request) {
	var data ConfirmUserRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		BadRequestError(err.Error()).Render(w, r)
		return
	}

	// Look up the user id associated with the passcode
	passcode := string(data.Passcode)
	userID, err := app.PasscodeStore.Find(passcode)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	}

	if userID != data.UserID {
		BadRequestError("Invalid passcode").Render(w, r)
		return
	}

	// Retrieve the user record
	user, err := app.UserStore.Find(userID)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	} else if user == nil {
		NotFoundError().Render(w, r)
		return
	}

	// Create the user session
	sessionID, err := app.SessionStore.Create(user.Id)
	if err != nil {
		panic(err)
	}

	// Create the session token
	sessionTokenSecret := []byte(app.Config.SessionToken.Secret)
	sessionClaims := session.NewSessionClaims(sessionID, app.Config)
	sessionToken, err := sessionClaims.Sign(sessionTokenSecret)
	if err != nil {
		panic(err)
	}

	// Set the session token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     app.Config.SessionCookie.Name,
		Value:    sessionToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   app.Config.SessionCookie.Secure,
		SameSite: http.SameSiteLaxMode,
	})

	// Create the access token
	accessTokenSecret := []byte(app.Config.AccessToken.Secret)
	identityClaims := identity.NewIdentityClaims(user.Id, sessionClaims, app.Config)
	accessToken, err := identityClaims.Sign(accessTokenSecret)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(ConfirmUserResponse{
		User: &User{
			Id: user.Id,
		},
		AccessToken: accessToken,
	})
}
