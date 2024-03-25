package main

import (
	"github.com/everywan/demo-server-go/commons/app"
)

func main() {
	app := app.New()
	// e := gin.New()

	// httpBundle := rest.New()
	// httpBundle.Run()
	app.AddBundle()
}
