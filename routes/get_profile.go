package routes

import (
	"encoding/json"
	"net/http"

	"github.com/birdbox/authnz/session"
)

type GetProfileResponse struct {
	*User       `json:"user"`
	AccessToken string `json:"access_token"`
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	// Read the session cookie from the request
	sessionCookie, err := r.Cookie(app.Config.Session.Name)
	if err != nil {
		panic(err)
	}

	// Parse the session cookie
	sessionClaims, err := session.Parse(sessionCookie.Value, []byte(app.Config.Server.Secret))
	if err != nil {
		panic(err)
	}

	// Look up the user ID from the session store
	sessionID := sessionClaims.Subject
	userID, err := app.SessionStore.Find(sessionID)
	if err != nil {
		panic(err)
	}

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
		AccessToken: "",
	})

}
