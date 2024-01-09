package main

import (
	"github.com/AaronSaikovski/golotteryservice/app"
	_ "github.com/AaronSaikovski/golotteryservice/docs"
)

// @title golotteryservice
// @version 0.1
// @description A random lottery generator API using Go & Fiber
// @contact.name Aaron Saikovski
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {

	// setup and run app
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
