package controllers

import (
	"encoding/json"
	"net/http"
	"rest-secure/models"
	u "rest-secure/utils"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// CreatePerson is the handler function for addind a person to the database
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uuid.UUID)
	person := &models.Person{}
	err := json.NewDecoder(r.Body).Decode(person)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	person.UserID = user
	resp := person.Create(App.DB)
	u.Respond(w, resp)
}

// GetPerson is the handler function for adding a specific person to the database
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}
	user := r.Context().Value("user").(uuid.UUID)
	data := models.GetPerson(user, id, App.DB)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetPersonsFor is the handler function for fetching current user's persons
func GetPersonsFor(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uuid.UUID)
	data := models.GetPersons(user, App.DB)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
