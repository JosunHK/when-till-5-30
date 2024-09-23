package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	eMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3030"
	}
	PORT = ":" + PORT

	e := echo.New()
	e.Use(eMiddleware.Recover())
	e.Use(eMiddleware.Logger())
	e.Pre(eMiddleware.RemoveTrailingSlashWithConfig(
		eMiddleware.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		},
	))

	e.Static("/", "public")

	e.Logger.Fatal(e.Start(PORT))
}
