package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/handler/homehandler"
	"github.com/uszebr/thegamem/internal/config"
	"github.com/uszebr/thegamem/internal/logger/loggerinit"
)

func main() {
	fmt.Println("Starting Thegamem..")
	// config path must be stored in env var CONFIG_PATH
	sv := config.MustLoad()
	//logger
	loggerinit.MustInitLogger(sv.Env)

	slog.Info("Server initialized with: ", "ENV", sv.Env, "Port", sv.AppPort)

	app := echo.New()
	public := app.Group("")

	public.Static("/static", "static")

	homeHandler := homehandler.HomeHandler{}
	public.GET("/", homeHandler.HandleShow)

	log.Fatal(app.Start(":" + sv.AppPort))
}
