package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/birdbox/authnz/identity"
	"github.com/birdbox/authnz/session"
)

type ConfirmUserRequest struct {
	Passcode string `json:"passcode"`
}

type ConfirmUserResponse struct {
	*User       `json:"user"`
	AccessToken string `json:"access_token"`
}

func confirmUser(w http.ResponseWriter, r *http.Request) {
	// Get the user-provided passcode from the request body and use it to
	// to look up the user id.
	var data ConfirmUserRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		BadRequestError(err.Error()).Render(w, r)
		return
	}

	// Validate that request body contains a passcode
	if data.Passcode == "" {
		BadRequestError("Missing passcode").Render(w, r)
		return
	}

	// Look up the user id associated with the passcode
	expectedID, err := app.PasscodeStore.Find(data.Passcode)
	if err != nil {
		panic(err)
	}

	if expectedID == 0 {
		UnauthorizedError("Invalid passcode").Render(w, r)
		return
	}

	// Get the authorization token from the request header
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		UnauthorizedError("Missing authentication token").Render(w, r)
		return
	}

	// Remove the "Bearer " prefix from the token and parse the claims
	claims, err := identity.Parse(bearerToken[7:], []byte(app.Config.Server.Secret))
	if err != nil {
		UnauthorizedError("Invalid authentication token").Render(w, r)
		return
	}

	actualID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		UnauthorizedError("Invalid authentication token").Render(w, r)
		return
	}

	if expectedID != actualID {
		UnauthorizedError("Invalid authentication token").Render(w, r)
		return
	}

	// Retrieve the user record
	user, err := app.UserStore.Find(expectedID)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	} else if user == nil {
		NotFoundError().Render(w, r)
		return
	}

	// Create the user session
	sessionID, err := app.SessionStore.Create(user.ID)
	if err != nil {
		panic(err)
	}

	secret := []byte(app.Config.Server.Secret)

	// Create the session token
	sessionClaims := session.NewSessionClaims(sessionID, app.Config)
	sessionToken, err := sessionClaims.Sign(secret)
	if err != nil {
		panic(err)
	}

	// Set the session token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     app.Config.Session.Name,
		Value:    sessionToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   app.Config.Session.Secure,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().UTC().Add(time.Duration(app.Config.Session.MaxAge) * time.Second),
	})

	// Create the access token
	identityClaims := identity.NewIdentityClaims(user.ID, sessionClaims, app.Config)
	accessToken, err := identityClaims.Sign(secret)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(ConfirmUserResponse{
		User: &User{
			ID: user.ID,
		},
		AccessToken: accessToken,
	})
}
