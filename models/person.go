package models

import (
	"fmt"
	u "rest-secure/utils"

	"github.com/jinzhu/gorm"
)

// Person struct, as stored in the database
type Person struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	UserID    uint   `json:"user_id"`
}


// Validate checks the required parameters sent through the http request body
// returns message and true if the requirement is met
func (person *Person) Validate() (map[string]interface{}, bool) {
	if person.FirstName == "" {
		return u.Message(false, "Person's first name should be on the payload"), false
	}
	if person.LastName == "" {
		return u.Message(false, "Person's last name should be on the payload"), false
	}
	if person.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}
	if person.Email == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}
	if person.UserID <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

// Create adds a new person to the database
func (person *Person) Create() map[string]interface{} {
	if resp, ok := person.Validate(); !ok {
		return resp
	}
	GetDB().Create(person)
	resp := u.Message(true, "success")
	resp["person"] = person
	return resp
}

// GetPerson returns a single person, if present, that matches provided criteria
func GetPerson(user, id uint) *Person {
	person := &Person{}
	err := GetDB().Where(&Person{UserID: user}).First(&person, id).Error
	if err != nil {
		return nil
	}
	return person
}

// GetPerson returns an array of persons for current user
func GetPersons(user uint) []*Person {
	persons := make([]*Person, 0)
	err := GetDB().Where(&Person{UserID: user}).Find(&persons).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return persons
}
