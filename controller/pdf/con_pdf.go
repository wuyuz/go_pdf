package pdf

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"pdfUtil/config"
	"pdfUtil/controller"
	"pdfUtil/lib"
	"pdfUtil/services"
)

type PdfController struct {
	pdfService *services.PdfService
}

func NewPdfController(pdfService *services.PdfService) controller.PdfInterface {
	return &PdfController{pdfService}
}

func (p *PdfController)ReaderPdf(c *fiber.Ctx) error {
	// Get first file from form field "updatePdf":
	file, err := c.FormFile("updatePdf")
	if err != nil {
		return c.JSON(fiber.Map{
			"code": "602",
			"msg":  "Invalid file",
		})
	}

	inpuFile := config.SAVEFILE +  file.Filename
	outFile := config.SENDFILE + file.Filename
	defer func() {
		if _, err := os.Stat(inpuFile); err==nil {
			os.Remove(inpuFile)
		}
		if _, err := os.Stat(outFile); err==nil {
			os.Remove(outFile)
		}
	}()

	// Save file to root director
	err = c.SaveFile(file, inpuFile)
	if err != nil {
		return c.JSON(fiber.Map{
			"code": "603",
			"msg":  "Save input file error",
		})
	}

	p.pdfService.AlertAttr(inpuFile,outFile,config.COVER)

	if err = p.pdfService.AddWatermarks(false); err != nil {
		lib.CheckErr(err,"pdf merge error")
	}

	return c.SendFile(outFile,false)
}
