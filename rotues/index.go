package rotues

import (
	"github.com/gofiber/fiber/v2"
	"pdfUtil/rotues/pdf"
)

func  InitRoutes(app *fiber.App) {
	// 一级路由
	api := app.Group("/api/v1")

	// 二级路由
	pdf_router := api.Group("/pdf")
	pdf.PdfUpload(pdf_router)
}
