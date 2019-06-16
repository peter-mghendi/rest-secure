package controllers

import (
	"encoding/json"
	"net/http"
	"rest-secure/models"
	u "rest-secure/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var CreatePerson = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint) // Grab the id of the user that send the request
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

var GetPersonsFor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}
	data := models.GetPersons(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
