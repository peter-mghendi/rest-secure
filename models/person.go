package models

import (
	"fmt"
	u "rest-secure/utils"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Person struct, as stored in the database
type Person struct {
	Base
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	UserID    uuid.UUID `json:"user_id"`
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

	return u.Message(true, "success"), true
}

// Create adds a new person to the database
func (person *Person) Create(db *gorm.DB) map[string]interface{} {
	if resp, ok := person.Validate(); !ok {
		return resp
	}
	db.Create(person)
	resp := u.Message(true, "success")
	resp["person"] = person
	return resp
}

// GetPerson returns a single person, if present, that matches provided criteria
func GetPerson(user, id uuid.UUID, db *gorm.DB) *Person {
	person := &Person{}
	err := db.Where(&Person{Base: Base{ID: id}, UserID: user}).First(person).Error
	if err != nil {
		return nil
	}
	return person
}

// GetPersons returns an array of persons for current user
func GetPersons(user uuid.UUID, db *gorm.DB) []*Person {
	persons := make([]*Person, 0)
	err := db.Where(&Person{UserID: user}).Find(&persons).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return persons
}
