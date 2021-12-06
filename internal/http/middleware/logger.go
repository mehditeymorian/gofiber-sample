package middleware

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Log struct {
	Type      string    `json:"type,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Path      string    `json:"path,omitempty"`
	Data      string    `json:"data,omitempty"`
}

func NewLogger(ctx *fiber.Ctx) error {

	logData := Log{
		Type:      "request",
		Timestamp: time.Now(),
		Path:      ctx.Request().URI().String(),
		Data:      string(ctx.Request().Body()),
	}

	bytes, _ := json.Marshal(logData)

	log.Println(string(bytes))

	return ctx.Next()
}
