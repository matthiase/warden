package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/matthiase/warden/identity"
	"github.com/matthiase/warden/session"
)

type GetProfileResponse struct {
	*User       `json:"user"`
	AccessToken string `json:"access_token"`
}

func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	var identityToken string

	if identityTokenCookie, err := r.Cookie(app.Config.Session.Name + "_it"); err == nil {
		identityToken = identityTokenCookie.Value
	}

	if identityToken == "" {
		// If the identity token cookie is missing, try to generate a new one using the session
		// cookie. If the session cookie is missing, then the user is not authenticated. Since
		// this endpoint is accessible to unauthenticated users, return an empty response instead
		// of an error.
		sessionTokenCookie, err := r.Cookie(app.Config.Session.Name + "_st")
		if err != nil {
			UnauthorizedError("Missing session token").Render(w, r)
			return
		}

		// Parse the session token to get the session id. If the session token is invalid,
		// return an error.
		sessionClaims, err := session.Parse(sessionTokenCookie.Value, []byte(app.Config.Server.Secret))
		if err != nil {
			UnauthorizedError("Invalid session token").Render(w, r)
			return
		}

		// Look up the user id associated with the session id
		userID, err := app.SessionStore.Find(sessionClaims.Subject)
		if err != nil {
			panic(err)
		}

		// Fetch the user record from the database
		user, err := app.UserStore.Find(userID)
		if err != nil {
			UnauthorizedError("Invalid session token").Render(w, r)
			return
		}

		// Generate a new identity token for the user
		identityClaims := identity.NewIdentityClaims(sessionClaims.Subject, user, app.Config)
		identityToken, err = identityClaims.Sign([]byte(app.Config.Server.Secret))
		if err != nil {
			panic(err)
		}

		// Set the identity token cookie
		http.SetCookie(w, &http.Cookie{
			Name:     app.Config.Session.Name + "_it",
			Value:    identityToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   app.Config.Session.Secure,
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Now().UTC().Add(3600 * time.Second),
		})
	}

	// At this point, the identity token has either been retrieved from the cookie or
	// generated using the session token. Parse the identity token to get the user id.
	identityClaims, err := identity.Parse(identityToken, []byte(app.Config.Server.Secret))
	if err != nil {
		UnauthorizedError("Invalid identity token").Render(w, r)
		return
	}

	userID := identityClaims.Subject

	// Get the user from the database
	user, err := app.UserStore.Find(userID)
	if err != nil {
		panic(err)
	}

	if user == nil {
		NotFoundError().Render(w, r)
		return
	}

	json.NewEncoder(w).Encode(GetProfileResponse{
		User: &User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
		AccessToken: identityToken,
	})
}
