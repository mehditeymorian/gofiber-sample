package model

import (
	"time"
)

type Person struct {
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Age         int       `json:"age,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewPerson(
	name string,
	email string,
	phoneNumber string,
	age int,
	createdAt time.Time,
) *Person {
	return &Person{
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Age:         age,
		CreatedAt:   createdAt,
	}
}
