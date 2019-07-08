package main

import (
	"fmt"
	"net/http"
	"os"
	"rest-secure/app"
	"rest-secure/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) //attach JWT auth middleware
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/persons", controllers.GetPersonsFor).Methods("GET")

	router.HandleFunc("/api/persons/new", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/api/persons/{id}", controllers.GetPerson).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Serving on localhost:%v\n", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
