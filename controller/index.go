package controller

import "github.com/gofiber/fiber/v2"

type PdfInterface interface {
	ReaderPdf(c *fiber.Ctx) error
}
