package controllers

import (
	"encoding/json"
	"net/http"
	"rest-secure/models"
	u "rest-secure/utils"
)

// CreateAccount is the handler function for adding a new account into the databsse
func CreateAccount (w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create(App.DB)
	u.Respond(w, resp)
}

// Authenticate is the handler function for aurhorizing user login
func Authenticate (w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password, App.DB)
	u.Respond(w, resp)
}
