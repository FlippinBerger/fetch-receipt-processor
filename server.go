package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// I like to register my routes in a separate file because it can get unruly
	// in a large API
	registerRoutes(e)

	// Default testing port kept here because it doesn't particularly matter for
	// this take home. Would update the whole thing to use TLS over 443 or similar in prod
	e.Logger.Fatal(e.Start(":1323"))
}
