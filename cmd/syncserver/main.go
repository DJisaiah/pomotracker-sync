package main

import (
	"fmt"
	"github.com/DJisaiah/pomotracker-sync/internal/server"
)

func main(){
	fmt.Println("starting server")
	server.Start()
}
