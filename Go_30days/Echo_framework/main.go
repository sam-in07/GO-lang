package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	dsn := "host=localhost port=5432 user=postgres password=admin dbname=newDb sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect DB:", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE,
		age INT
	);
	`

	if _, err := db.Exec(createTable); err != nil {
		log.Fatal("Failed to create table:", err)
	}
    e.POST("/users",func(c echo.Context) error {
		u := new(User)
		if err:=c.Bind(u); err!=nil {
			return c.JSON(http.StatusBadRequest , map[string]string{"error" : "Invalid request Body "})
		}
		var id int 
        err:= db.QueryRow(
			""
		)

	})
	e.Logger.Fatal(e.Start(":8090"))

}
