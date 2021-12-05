package request

import (
	"testing"
)

func TestAge(t *testing.T) {
	tests := []struct {
		name   string
		ok     bool
		person Person
	}{
		{
			name: "underage",
			ok:   false,
			person: Person{
				Age: -2,
			},
		},
		{
			name: "overage",
			ok:   false,
			person: Person{
				Age: 125,
			},
		},
		{
			name: "correct age",
			ok:   true,
			person: Person{
				Age: 23,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.person.Validate()
			hasErr := err != nil

			if hasErr == test.ok {
				t.Errorf("failed: %v", err)
			}
		})
	}
}
