package main

import (
	"fmt"
	"os"
	"os/exec"
)

func getDBConfig() string {
	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	password := os.Getenv("DB_PASSWORD")

	CONNECT := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4&parseTime=true"

	return CONNECT
}

func executeInitialize() error {
	CONNECT := getDBConfig()
	fmt.Printf("result: %s", CONNECT)
	out, err := exec.Command("migrate", "-path", "migration/db/migration", "-database", CONNECT, "1", "up").Output()

	if err != nil {
		return err
	}
	fmt.Printf("result: %s", out)
	return nil
}

func main() {
	executeInitialize()
}
