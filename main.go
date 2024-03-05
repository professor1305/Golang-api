package main

import (
	"log"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"api/Model"
	
)

type Userdetail struct {
	ID          int    `json:"ID"`
	Address     string `json:"Address"`
	Education   string `json:"Education"`
	Gender      string `json:"Gender"`
	Department  string `json:"Department"`
}

func main() {
	app := fiber.New()
	
	userDetails := []Userdetail{
		{
			ID:          1,
			Address:     "123 Main St",
			Education:   "Bachelor's Degree",
			Gender:      "Male",
			Department:  "Engineering",
		},
		{
			ID:          2,
			Address:     "456 Elm St",
			Education:   "Master's Degree",
			Gender:      "Female",
			Department:  "Marketing",
		},
		{
			ID:          3,
			Address:     "789 Oak St",
			Education:   "PhD",
			Gender:      "Male",
			Department:  "Research",
		},
	}

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
	app.Get("/Usersdetails", func(c *fiber.Ctx) error {
    query := c.Query("ID")
    var result []Userdetail
    for _, user := range userDetails {
        if query == strconv.Itoa(user.ID) {
            result = append(result, user)
        }
    }
    if len(result) > 0 {
        return c.JSON(result)
    } else {
        return c.Status(fiber.StatusNotFound).JSON(map[string]string{"status": "not found"})
    }
})

	log.Fatal(app.Listen(":3000"))
}
