package app

import (
	"fmt"
	"log"
	"net/http"

	"rest-secure/controllers"
	"rest-secure/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" // init postgresql drivers
)

var err error

// App struct holds router and db references
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) initRoutes() {
	a.Router = mux.NewRouter()
	a.Router.Use(JwtAuthentication)
	a.Router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	a.Router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	a.Router.HandleFunc("/api/me/persons", controllers.GetPersonsFor).Methods("GET")
	a.Router.HandleFunc("/api/persons/new", controllers.CreatePerson).Methods("POST")
	a.Router.HandleFunc("/api/persons/{id}", controllers.GetPerson).Methods("GET")
}

func (a *App) initDB(dbHost, dbUser, dbName, dbPass, dbType string) {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPass)
	// fmt.Println(dbURI)

	a.DB, err = gorm.Open(dbType, dbURI)
	if err != nil {
		fmt.Print(err)
	}

	a.DB.Debug().AutoMigrate(&models.Account{}, &models.Person{})
}

// HACK avoid cyclic dependencies by "injecting" a into controllers package
func (a *App) initVars() {
	controllers.App.Router, controllers.App.DB = a.Router, a.DB
}

// Init sets up routes and database connection
func (a *App) Init(dbHost, dbUser, dbName, dbPass, dbType string) {
	a.initDB(dbHost, dbUser, dbName, dbPass, dbType)
	a.initRoutes()
	a.initVars()
}

// Run serves the API on a specified port
func (a *App) Run(port string) {
	fmt.Printf("Serving on localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}
