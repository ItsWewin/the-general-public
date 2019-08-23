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
	"os"
)

var dbConf = config.DBConfig

func main() {
	action := flag.String("action", "up", "Set your action up/down")
	steps := flag.Int("steps", 1, "Set the steps，when you set the action 'down', you should set the steps")
	flag.Parse()

	//dataSourceName  "root:wewin123@tcp(127.0.0.1:3306)/general_public?multiStatements=true"
	dataSourceNmae := fmt.Sprintf("%v:%v@%v(%v:%v)/%v?multiStatements=true",
								dbConf.UserName, dbConf.PassWord, dbConf.Protocol, dbConf.Host, dbConf.Port, dbConf.DBName)
	db, err := sql.Open("mysql", dataSourceNmae)
	if err != nil {
		log.Fatal("Some error occurred when connect mysql, error: %v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Some error occurred when instance, error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%v", dbConf.MigrationsFilePath),
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Some error occurred when database instance, error: %v", err)
	}

	switch *action {
	case "up":
		if err := m.Up(); err != nil {
			if err.Error() == "no change" {
				log.Println("All change have been migrated")
				os.Exit(0)
			} else {
				log.Fatalf("An error occurred when sysncing the database, error: %v", err)
			}
		}
	case "down":
		if err := m.Steps(-*steps); err != nil {
			if err.Error() == "file does not exist" {
				log.Fatal("Do not have any migration")
			} else {
				log.Fatalf("An error occurred when sysncing the database, error: %v", err)
			}
		}
	default:
		fmt.Printf("Unkown action, action must be 'up' or 'down'")
	}
}