package main

import (
	"CALC_GO/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}