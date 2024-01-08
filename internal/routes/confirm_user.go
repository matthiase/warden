package routes

import (
	"encoding/json"
	"net/http"
)

func confirmUser(w http.ResponseWriter, r *http.Request) {
	var data ConfirmUserRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		BadRequestError(err.Error()).Render(w, r)
		return
	}

	account, err := application.AccountStore.Find(data.Id)
	if err != nil {
		ApplicationError(err.Error()).Render(w, r)
		return
	} else if account == nil {
		NotFoundError().Render(w, r)
		return
	}

	// TODO: validate the OTP

	// TODO: generate an access token

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ConfirmUserResponse{
		User: &User{
			Id:    account.Id,
			Email: account.Email,
			Name:  account.Name,
		},
		AccessToken: "qwertyuiopasdfghjklzxcvbnm",
	})
}
