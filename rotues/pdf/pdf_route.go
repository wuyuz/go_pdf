package pdf

import (
	"github.com/gofiber/fiber/v2"
	"pdfUtil/controller/pdf"
	"pdfUtil/services"
)
func PdfUpload(router fiber.Router) {
	//  在刚进入服务的时候会初始化service的实现对象传入db
	p := pdf.NewPdfController(services.NewPdfService())

	router.Post(
		"/upload",
		p.ReaderPdf,
	)
}
