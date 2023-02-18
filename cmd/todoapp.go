package main

import (
	"fmt"
	"html/template"
	"log"

	"github.com/ashwins93/wutplans/db"
	httpService "github.com/ashwins93/wutplans/http"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbUrl := "file:./dev.db?_fk=1"
	fmt.Println("Connecting to", dbUrl)
	conn, err := sqlx.Connect("sqlite3", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	db := db.New(conn)

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Static("/", "public")

	service := httpService.New(db)
	service.SetupRoutes(e)

	templates := template.Must(template.ParseGlob("http/html/*.html"))
	e.Renderer = service.GetTemplateRenderer(templates)

	log.Fatal(e.Start(":3000"))
}
