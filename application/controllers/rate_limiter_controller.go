package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type RateLimiterController struct{}

func NewRateLimiterController() *RateLimiterController {
	return &RateLimiterController{}
}

func (c *RateLimiterController) GetController(ctx *fiber.Ctx) error {
	log.Printf("Processo encerrado com sucesso")

	return ctx.Status(200).JSON("Success")
}
