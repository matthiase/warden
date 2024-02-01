package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/birdbox/authnz/identity"
	"github.com/birdbox/authnz/session"
)

type ConfirmUserRequest struct {
	Passcode string `json:"passcode"`
}

func confirmUser(w http.ResponseWriter, r *http.Request) {
	// Get the user-provided passcode from the request body and use it to
	// to look up the associated user id.
	var data ConfirmUserRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		BadRequestError(err.Error()).Render(w, r)
		return
	}

	if data.Passcode == "" {
		BadRequestError("Missing passcode").Render(w, r)
		return
	}

	expectedUserID, err := app.PasscodeStore.Find(data.Passcode)
	if err != nil {
		UnauthorizedError("Invalid passcode").Render(w, r)
		return
	}

	// Parse the value of the verification token cookie to get the user id.
	// If the user id in the verification token matches the expected user id,
	// then the user has successfully confirmed their identity.
	verificationToken, err := r.Cookie(app.Config.Session.Name + "_vt")
	if err != nil {
		UnauthorizedError("Missing verification token").Render(w, r)
		return
	}

	// Parse the verification token to get the user id
	verificationClaims, err := identity.Parse(verificationToken.Value, []byte(app.Config.Server.Secret))
	if err != nil {
		UnauthorizedError("Invalid authentication token").Render(w, r)
		return
	}

	providedUserID := verificationClaims.Subject

	if expectedUserID != providedUserID {
		UnauthorizedError("Invalid authentication token").Render(w, r)
		return
	}

	// Fetch the user record from the database and create a new session.
	user, err := app.UserStore.Find(providedUserID)
	if err != nil {
		UnauthorizedError("Invalid authentication token").Render(w, r)
		return
	}

	// Print the user's name to the console
	log.Printf("User %s %s (%s) has confirmed their identity", user.FirstName, user.LastName, user.Email)

	sessionID, err := app.SessionStore.Create(user.ID)
	if err != nil {
		panic(err)
	}

	// Generate the session and identity tokens and set them as http-only
	// cookies in the response.
	secret := []byte(app.Config.Server.Secret)

	// The session cookie will be used to refresh expired identity tokens
	sessionClaims := session.NewSessionClaims(sessionID, app.Config)
	sessionToken, err := sessionClaims.Sign(secret)
	if err != nil {
		panic(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     app.Config.Session.Name + "_st",
		Value:    sessionToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   app.Config.Session.Secure,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().UTC().Add(time.Duration(app.Config.Session.MaxAge) * time.Second),
	})

	// The identity token will be used to authenticate requests
	identityClaims := identity.NewIdentityClaims(sessionID, user, app.Config)
	identityToken, err := identityClaims.Sign(secret)
	if err != nil {
		panic(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     app.Config.Session.Name + "_it",
		Value:    identityToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   app.Config.Session.Secure,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().UTC().Add(3600 * time.Second),
	})

	// Remove the verification token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     app.Config.Session.Name + "_vt",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   app.Config.Session.Secure,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().UTC().Add(-1 * time.Second),
	})

	// Also return the verification token in the response body?

}
