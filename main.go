package main

import (
	"api/Model"
	"api/Services"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)



func main() {
	app := fiber.New()
	

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5500",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
	}))

	app.Get("/Users", Model.GetUserDeatails)
	app.Get("/User", Model.GetAllusers)
	app.Post("/Users", Model.CreateUser)
	app.Get("/Usersdetails", Services.GetUserDetailsByID)

	log.Fatal(app.Listen(":3000"))
}
