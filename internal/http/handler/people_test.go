package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mehditeymorian/gofiber-sample/internal/cache"
	"github.com/mehditeymorian/gofiber-sample/internal/http/request"
	"github.com/mehditeymorian/gofiber-sample/internal/model"
)

type testSuit struct {
	app   *fiber.App
	cache cache.Local
}

func Test_People(t *testing.T) {
	cache := cache.New()
	app := fiber.New()
	v1 := app.Group("/v1")
	People{Cache: cache}.Register(v1)

	test := testSuit{
		app:   app,
		cache: cache,
	}

	t.Run("Set", test.SetTest)
	t.Run("Get", test.GetTest)
	t.Run("Del", test.DelTest)

}

func (suit testSuit) SetTest(t *testing.T) {
	tests := []struct {
		name         string
		route        string
		expectedCode int
		data         request.Person
	}{
		{
			name:         "correct request",
			route:        "/v1/people",
			expectedCode: http.StatusOK,
			data: request.Person{
				Name:        "name",
				Email:       "name@gmail.com",
				PhoneNumber: "0123",
				Age:         22,
			},
		},
		{
			name:         "underage request",
			route:        "/v1/people",
			expectedCode: http.StatusBadRequest,
			data: request.Person{
				Name:        "name",
				Email:       "name@gmail.com",
				PhoneNumber: "0123",
				Age:         -15,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			data, _ := json.Marshal(test.data)
			req := httptest.NewRequest("POST", test.route, bytes.NewReader(data))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := suit.app.Test(req, 1)

			assertResult(test.expectedCode, resp, t)
		})
	}
}

func (suit testSuit) GetTest(t *testing.T) {

	person := model.NewPerson("test-name", "test@gmail.com", "123", 20, time.Now())

	key, _ := suit.cache.Set("", *person)

	tests := []struct {
		name         string
		route        string
		expectedCode int
	}{
		{
			name:         "correct request",
			route:        "/v1/people/" + key,
			expectedCode: http.StatusOK,
		},
		{
			name:         "no content request",
			route:        "/v1/people/non-existing-key",
			expectedCode: http.StatusNoContent,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", test.route, nil)

			resp, _ := suit.app.Test(req, 1)

			assertResult(test.expectedCode, resp, t)
		})
	}
}

func (suit testSuit) DelTest(t *testing.T) {

	person := model.NewPerson("name-of-user-who-is-about-to-be-deleted", "test@gmail.com", "123", 40, time.Now())

	key, _ := suit.cache.Set("", *person)

	tests := []struct {
		name         string
		route        string
		expectedCode int
	}{
		{
			name:         "correct request",
			route:        "/v1/people/" + key,
			expectedCode: http.StatusNoContent,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", test.route, nil)

			resp, _ := suit.app.Test(req, 1)

			assertResult(test.expectedCode, resp, t)
		})
	}

}

func assertResult(expectedCode int, resp *http.Response, t *testing.T) {
	if resp.StatusCode != expectedCode {
		t.Logf("expectedCode: %d actualCode %d", expectedCode, resp.StatusCode)
		responseData, _ := ioutil.ReadAll(resp.Body)
		t.Error(string(responseData))
	}
}
