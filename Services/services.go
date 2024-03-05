package Services

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Userdetail struct {
	ID          int    `json:"ID"`
	Username    string `json:"Username"`
	Email       string `json:"Email"`
	Address     string `json:"Address"`
	Education   string `json:"Education"`
	Gender      string `json:"Gender"`
	Department  string `json:"Department"`
}
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "postgres://myuser:mypassword@0.0.0.0:5432/mydatabase?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetUserDetailsByID(c *fiber.Ctx) error {
    queryID := c.Query("ID")

    dialect := goqu.Dialect("postgres")
    query := dialect.From("usertable").Select("usertable.id", "usertable.username", "usertable.email", "userdetail.address", "userdetail.education", "userdetail.gender", "userdetail.department").
        Join(
            goqu.I("userdetail").As("userdetail"),
            goqu.On(goqu.Ex{"usertable.id": goqu.I("userdetail.id")}),
        ).
        Where(goqu.Ex{"usertable.id": queryID})

    sqlString, _, err := query.ToSQL()
    if err != nil {
        return err
    }
	fmt.Println(sqlString)
    rows, err := Db.Query(sqlString)
    if err != nil {
        return err
    }
    defer rows.Close()

    var result []Userdetail
    for rows.Next() {
        var user Userdetail
        if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Address, &user.Education, &user.Gender, &user.Department); err != nil {
            return err
        }
        result = append(result, user)
    }

    if len(result) > 0 {
        return c.JSON(result)
    } else {
        return c.Status(fiber.StatusNotFound).JSON(map[string]string{"status": "not found"})
    }
}

