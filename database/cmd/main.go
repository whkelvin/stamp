package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("dotenv failed")
	}

	fmt.Println(os.Getenv("POSTGRES_CONNECTION_STRING"))
	fmt.Println("hello")

	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_CONNECTION_STRING"))
	if err != nil {
		fmt.Println(err.Error())
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "INSERT INTO post (link, title, description) VALUES ('this a link', 'this is a title', 'this is a description');")

	if err != nil {
		fmt.Println(err.Error())
	}
}
