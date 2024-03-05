package Model

import (
	"github.com/gofiber/fiber/v2"
	 "database/sql"
	_ "github.com/lib/pq"
	"github.com/doug-martin/goqu/v9"
)

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
}
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "postgres://myuser:mypassword@0.0.0.0:5432/mydatabase?sslmode=disable")
	if err != nil {
		panic(err)
	}
}


func GetUserDeatails(c *fiber.Ctx) error {
		//Dialect := goqu.Dialect("postgres")
		var users []User
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
	}
func GetAllusers(c *fiber.Ctx) error {
    
    Dialect := goqu.Dialect("postgres")
    query := Dialect.From("usertable").Select("*")

    // Execute the query and retrieve the results
    sqlString, _, err := query.ToSQL()
    if err != nil {
        return err
    }

    // Execute the query and retrieve the results
    rows, err := Db.Query(sqlString)
    if err != nil {
        return err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
            return err
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return err
    }

    // Return the list of users as a JSON response
    return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {

	var newUser User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Insert the new user into the database
	dialect := goqu.Dialect("postgres")
	query := dialect.Insert("usertable").Rows(newUser)

	sqlString, _, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = Db.Exec(sqlString)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(newUser)
}

