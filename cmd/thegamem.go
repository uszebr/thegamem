package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/handler/homehandler"
	"github.com/uszebr/thegamem/handler/loginhandler"
	"github.com/uszebr/thegamem/handler/newgamehandler"
	"github.com/uszebr/thegamem/internal/authservice"
	"github.com/uszebr/thegamem/internal/config"
	"github.com/uszebr/thegamem/internal/jwtservice"
	"github.com/uszebr/thegamem/internal/logger/loggerinit"
	"github.com/uszebr/thegamem/internal/middleware/usermiddleware"
	"github.com/uszebr/thegamem/internal/supa"
	"github.com/uszebr/thegamem/play/model/modelfactory"
	"github.com/uszebr/thegamem/play/usergames"
)

func main() {
	fmt.Println("Starting Thegamem..")
	// config path must be stored in env var CONFIG_PATH
	sv := config.MustLoad()
	//logger
	loggerinit.MustInitLogger(sv.Env)

	slog.Info("Server initialized with: ", "ENV", sv.Env, "Port", sv.AppPort)

	supaAuthWrapperClient := supa.GetSupaAuth()
	jwtClient := jwtservice.JwtService{}
	authservice := authservice.New(supaAuthWrapperClient, jwtClient)
	userMiddleware := usermiddleware.New(authservice)
	app := echo.New()
	public := app.Group("")
	public.Use(userMiddleware.GetUserForPublic)
	public.Static("/static", "static")

	homeHandler := homehandler.HomeHandler{}
	public.GET("/", homeHandler.HandleShow)

	loginHandler := loginhandler.New(authservice)

	public.GET("/login", loginHandler.HandleShow)
	public.POST("/login", loginHandler.HandlePost)

	//loggedin urls
	loggedIn := app.Group("")
	loggedIn.Use(userMiddleware.LoggedIn)
	loggedIn.POST("/logout", loginHandler.LogoutPost)

	newGameHandler := newgamehandler.New(usergames.GetUserGames(), modelfactory.GetModelFactory())

	loggedIn.GET("/newgame", newGameHandler.HandleShow)
	loggedIn.POST("/newgame", newGameHandler.HandlePost)

	loggedIn.GET("/game", newGameHandler.HandleExistingGame)
	loggedIn.POST("/addboard", newGameHandler.HandleAddBoardPost)

	loggedIn.GET("/boards/:id", newGameHandler.HandleBoard)
	loggedIn.POST("/boardroundsforplayer", newGameHandler.HandleRoundsForPlayerPost)

	loggedIn.GET("/board/:boardId/round/:roundId", newGameHandler.HandleRound)

	log.Fatal(app.Start(":" + sv.AppPort))
}
