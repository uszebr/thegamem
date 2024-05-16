package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/handler/homehandler"
)

func main() {
	fmt.Println("Starting Thegamem server..")
	app := echo.New()
	public := app.Group("")

	public.Static("/static", "static")

	homeHandler := homehandler.HomeHandler{}
	public.GET("/", homeHandler.HandleShow)

	log.Fatal(app.Start(":8080"))
}
