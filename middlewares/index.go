package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)


func InitMiddleWares(app *fiber.App) {
	logFile, err := os.OpenFile("./logs/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Chdir()
	if err != nil {
		panic(err)
	}

	app.Use(logger.New(logger.Config{
		Format:     "#${pid} - ${time} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006/01/02 15:04:05",
		TimeZone:   "Local",
		Output: logFile,
	}))

}
