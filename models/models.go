package models

import "time"

// Customer schema of the customer table
type Customer struct {
    ID	int64	`json:"id"`
    FirstName	string	`json:"firstname"`
    LastName	string	`json:"lastname"`
	BirthDate	time.Time	`json:"birthdate"`
	Gender	string	`json:"gender"`
	Email	string	`json:"email"`
	Address	string	`json:"address"`
}