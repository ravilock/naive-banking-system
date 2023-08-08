package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/ravilock/naive-banking-system/internal/config"
	"github.com/ravilock/naive-banking-system/internal/db"
	"github.com/spf13/viper"
)

type Server interface {
	http.Handler
	Start()
}

type server struct {
	*echo.Echo
	db *db.DB
}

func (s *server) Start() {
	addr := fmt.Sprintf(":%d", viper.GetInt("port"))
	s.Echo.Logger.Fatal(s.Echo.Start(addr))
}

func NewServer() (Server, error) {
	dbInstance, err := db.New()
	if err != nil {
		return nil, err
	}

	if err := dbInstance.Pool.Ping(); err != nil {
		return nil, err
	}

	return newServerFromDB(dbInstance, false)
}

func newServerFromDB(dbInstance *db.DB, testing bool) (Server, error) {
	// Echo instance
	e := echo.New()
	e.HideBanner = true

	server := &server{
		Echo: e,
		db:   dbInstance,
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	apiGroup := e.Group("/api")
	apiGroup.GET("/healthcheck", server.healthcheck)

	// Start server
	e.Logger.Fatal(e.Start(":9191"))
	return server, nil
}
