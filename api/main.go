package main

import (
	"./src/system/app"
	"flag"
	"github.com/joho/godotenv"
	"os"
)

var port string

func init() {
	//get port from terminal eg go run main.go -port=3000
	flag.StringVar(&port, "port", "8000", "Using the assigned port.")
	flag.Parse()

	//fetches port from confog file
	err := godotenv.Load("config.ini")
	if err != nil {
		panic(err)
	}
	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}
func main() {
	s := app.NewServer()
	s.Init(port)
	s.Start()
}
