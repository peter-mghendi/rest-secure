package models

import (
	"fmt"
	u "rest-secure/utils"

	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	UserID    uint   `json:"user_id"` //The user that this person belongs to
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
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

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (person *Person) Create() map[string]interface{} {
	if resp, ok := person.Validate(); !ok {
		return resp
	}
	GetDB().Create(person)
	resp := u.Message(true, "success")
	resp["person"] = person
	return resp
}

func GetPerson(id uint) *Person {
	person := &Person{}
	err := GetDB().Table("persons").Where("id = ?", id).First(person).Error
	if err != nil {
		return nil
	}
	return person
}

func GetPersons(user uint) []*Person {
	persons := make([]*Person, 0)
	err := GetDB().Table("persons").Where("user_id = ?", user).Find(&persons).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return persons
}
