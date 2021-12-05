package cache

import (
	"fmt"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/mehditeymorian/gofiber-sample/internal/model"
)

type Local struct {
	people map[string]model.Person
}

func New() Local {
	return Local{
		people: make(map[string]model.Person),
	}
}
func (l Local) Set(optionalKey string, person model.Person) (string, error) {
	key := optionalKey
	if key == "" {
		key = generateRandomKey()
	}

	l.people[key] = person

	return key, nil
}

func (l Local) Get(key string) (*model.Person, error) {

	person, found := l.people[key]

	if !found {
		return nil, fmt.Errorf("no value exists for %s", key)
	}

	return &person, nil
}

func (l Local) Del(key string) error {

	delete(l.people, key)

	return nil
}

func generateRandomKey() string {
	return utils.UUIDv4()
}
