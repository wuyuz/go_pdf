package main

import (
	"github.com/gofiber/fiber/v2"
	"pdfUtil/config"
	"pdfUtil/lib"
	"pdfUtil/middlewares"
	"pdfUtil/rotues"
)

func main(){
	// 通过Setting可以拦截服务器的各种抛错
	app := fiber.New() // 创建 fiber app引擎
	// 初始化中间件
	middlewares.InitMiddleWares(app)


	rotues.InitRoutes(app)

	app.Get("/test",func(c *fiber.Ctx) error {
		return c.SendFile("./test.html");})


	// 启动服务
	err := app.Listen(config.PORT)
	lib.CheckErr(err,"listening")
}