package main

import (
	"fmt"
	"github.com/ItsWewin/the-general-public/config"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	fileName := strings.Join(args, "_")
	timeStamp := time.Now().Unix()

	migrationPath := config.DBConfig.MigrationsFilePath

	fullUPFileName := fmt.Sprintf("%v/%v_%v.up.sql", migrationPath, timeStamp, fileName)
	fullDownFileName := fmt.Sprintf("%v/%v_%v.down.sql", migrationPath, timeStamp, fileName)

	f, err := os.Create(fullUPFileName)
	if err != nil {
		fmt.Printf("One error occurred when create file %v, error: %v", fullUPFileName, err)
	}
	defer f.Close()

	f2, err := os.Create(fullDownFileName)
	if err != nil {
		fmt.Printf("One error occurred when create file %v, error: %v", fullDownFileName, err)
	}
	defer f2.Close()

	fmt.Printf("Crete file: %v\n", fullUPFileName)
	fmt.Printf("Created file: %v\n", fullDownFileName)
}