package main

import (
	"github.com/golangee/log"
	"github.com/golangee/log/simple"
	"rochus/frontend/internal/app"
)

func main() {
	log.SetDefault(simple.PrintColored)
	a := app.NewApplication()
	a.Run()
}
