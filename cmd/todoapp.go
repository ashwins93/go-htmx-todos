package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	// "github.com/ashwins93/wutplans"
	"github.com/ashwins93/wutplans/db"
	"github.com/jmoiron/sqlx"
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

	/*todo, err := db.CreateTodo(context.Background(), &wutplans.Todo{
		Task: "Explore htmx",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Todo created", todo)
	*/

	todos, err := db.FindAllTodos(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i, todo := range todos {
		json, _ := json.MarshalIndent(todo, "", "  ")
		fmt.Printf("Todo %d: %v\n", i+1, string(json))
	}
}
