package controllers

import (
	"encoding/json"
	"net/http"
	"rest-secure/models"
	u "rest-secure/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePerson is the handler function for addind a person to the database
var CreatePerson = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	person := &models.Person{}
	err := json.NewDecoder(r.Body).Decode(person)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	person.UserID = user
	resp := person.Create()
	u.Respond(w, resp)
}

// GetPerson is the handler function for adding a specific person to the database
var GetPerson = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}
	user := r.Context().Value("user").(uint)
	data := models.GetPerson(user, uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetPersonsFor is the handler function for fetching current user's persons
var GetPersonsFor = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	data := models.GetPersons(user)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
