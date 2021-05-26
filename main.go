package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"mediatR/API"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	fmt.Println("Welcome to MediatR pattern")
	API.Startup()
}
