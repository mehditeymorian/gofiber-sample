package request

import (
	"fmt"
)

type Person struct {
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Age         int    `json:"age,omitempty"`
}

func (p Person) Validate() error {
	if p.Age < 0 {
		return fmt.Errorf("you are %d! come back when you are born", p.Age)
	} else if p.Age > 100 {
		return fmt.Errorf("kudos to you! how are you still alive? :)) sorry pal we have age limit")
	}

	return nil
}
