package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/ItsWewin/the-general-public/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
)

func main() {
	action := flag.String("action", "up", "select your action up/down")
	flag.Parse()

	db, err := sql.Open("mysql", "root:wewin123@tcp(127.0.0.1:3306)/general_public?multiStatements=true")
	if err != nil {
		log.Fatal("Some error occurred when connect mysql, error: %v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("Some error occurred when instance, error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%v", config.DBConfig.MigrationsFilePath),
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal("Some error occurred when database instance, error: %v", err)
	}

	switch *action {
	case "up":
		if err := m.Up(); err != nil {
			fmt.Println("up action")
			log.Fatal("An error occurred when sysncing the database, error: %v", err)
		}
	case "down":
		if err := m.Down(); err != nil {
			fmt.Println("down action")
			log.Fatal("An error occurred when sysncing the database, error: %v", err)
		}
	default:
		fmt.Printf("Unkown action, action must be 'up' or 'down'")
	}
}