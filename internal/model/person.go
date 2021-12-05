package model

import (
	"time"
)

type Person struct {
	Name        string
	Email       string
	PhoneNumber string
	Age         int
	CreatedAt   time.Time
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
