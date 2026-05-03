package main

import (
	"tic-tac-toe/internal/di"
)

func main() {
	app := di.BuildContainer()
	app.Run()
}
