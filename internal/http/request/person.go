package request

import (
	"fmt"
)

type Person struct {
	Name        string
	Email       string
	PhoneNumber string
	Age         int
}

func (p Person) Validate() error {
	if p.Age < 0 {
		return fmt.Errorf("you are %d! come back when you are born", p.Age)
	} else if p.Age > 100 {
		return fmt.Errorf("kudos to you! how are you still alive? :)) sorry pal we have age limit")
	}

	return nil
}
