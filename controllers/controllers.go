package controllers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Application struct {
	Router *mux.Router
	DB     *gorm.DB
}

var App *Application

func init() {
	App = &Application{}
}

