package main

import (
	"fmt"
	"log"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
}
type Userdetail struct {
	ID          int    `json:"ID"`
	Address     string `json:"Address"`
	Education   string `json:"Education"`
	Gender      string `json:"Gender"`
	Department  string `json:"Department"`
}

func main() {
	app := fiber.New()
	var users []User
	users = append(users, User{ID: 1, Username: "alice", Email: "alice@example.com"})
	users = append(users, User{ID: 2, Username: "bob", Email: "bob@example.com"})
	users = append(users, User{ID: 3, Username: "charlie", Email: "charlie@example.com"})
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

	app.Get("/Users", func(c *fiber.Ctx) error {
		query := c.Query("Username")
		if query != "" {
			var filterUsers []string
			if query == "Username" {
				for _, user := range users {
					filterUsers = append(filterUsers, user.Username)
				}
				return c.JSON(filterUsers)
			}
			if query == "Email" {
				for _, user := range users {
					filterUsers = append(filterUsers, user.Email)
				}
				return c.JSON(filterUsers)
			}
			if query == "" {
				return c.JSON(users)
			}
		}
		return c.SendStatus(fiber.StatusOK)
	})
	app.Get("/User", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})
	app.Post("/Users", func(c *fiber.Ctx) error {
		var newUser User
		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(400).JSON(err)
		}
		users = append(users, newUser)
		fmt.Println(users)
		return c.Status(200).JSON(fiber.Map{
			"message": "User data received",
		})
	})
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
