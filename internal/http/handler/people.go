package handler

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mehditeymorian/gofiber-sample/internal/cache"
	"github.com/mehditeymorian/gofiber-sample/internal/http/request"
	"github.com/mehditeymorian/gofiber-sample/internal/http/response"
	"github.com/mehditeymorian/gofiber-sample/internal/model"
)

type People struct {
	Cache cache.Local
}

func (p People) Register(ctx fiber.Router) {
	ctx.Get("/people", p.GetAll)
	ctx.Get("/people/:key", p.Get)
	ctx.Post("/people", p.Set)
	ctx.Delete("/people/:key", p.Del)
}

func (p People) GetAll(ctx *fiber.Ctx) error {

	people := p.Cache.GetAll()

	return ctx.JSON(people)
}

func (p People) Get(ctx *fiber.Ctx) error {
	key := ctx.Params("key")

	person, err := p.Cache.Get(key)

	if err != nil {
		return fiber.NewError(http.StatusNoContent, err.Error())
	}

	return ctx.JSON(person)
}

func (p People) Set(ctx *fiber.Ctx) error {
	key := ctx.Query("key")
	var input request.Person

	if err := ctx.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if err := input.Validate(); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	person := model.NewPerson(
		input.Name,
		input.Email,
		input.PhoneNumber,
		input.Age,
		time.Now(),
	)

	key, err := p.Cache.Set(key, *person)

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	keyResponse := response.NewKey(key)

	return ctx.JSON(keyResponse)
}

func (p People) Del(ctx *fiber.Ctx) error {
	key := ctx.Params("key")

	if err := p.Cache.Del(key); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(http.StatusNoContent)
}
