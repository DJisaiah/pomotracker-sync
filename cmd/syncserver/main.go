package main

import (
	"fmt"
	"log"
	"os"
	"github.com/DJisaiah/pomotracker-sync/internal/server"
	"github.com/joho/godotenv"
	"github.com/DJisaiah/pomotracker-sync/internal/db"	
)

func main(){
	if err := godotenv.Load(); err != nil{
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	q, err := db.InitializePool(dsn)
	if err != nil {
		log.Fatal("Error connecting to db")
	}
	fmt.Println("starting server")
	server.StartServer()
}
