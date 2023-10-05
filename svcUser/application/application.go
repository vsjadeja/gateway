package application

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"user/database"

	"github.com/gorilla/mux"
	log "github.com/vsjadeja/log"
)

// An Option configures a Server.
type Option interface {
	apply(*Application)
}

type Application struct {
	name   string
	port   string
	ctx    context.Context
	logger *log.Logger
	router *mux.Router
}

func New(appName string, port string, router *mux.Router) *Application {
	app := &Application{
		name:   appName,
		logger: log.L(),
		port:   port,
		router: router,
	}
	if app.name != `` {
		app.logger = app.logger.Named(app.name)
	}

	return app
}

// Initialize initializes the app with predefined configuration
func (a *Application) Initialize() error {
	a.ctx = context.Background()
	a.logger.SetLevel(log.DebugLevel)
	a.logger = a.logger.Named(a.name)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	database.Connect(dsn)
	database.Migrate()

	return nil
}

// Run the app on it's router
func (a *Application) Run() {
	srv := &http.Server{
		Addr:         a.port,
		Handler:      a.router,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := srv.ListenAndServe()
	if err == nil {
		a.logger.Info(a.ctx, "Server is running", "AppName", a.name, "address", a.port)
	} else {
		a.logger.Error(a.ctx, "Server is stopped", "AppName", a.name, "address", a.port, "error", err.Error())
	}
}

// Name returns the name of the Application.
func (app *Application) Name() string { return app.name }

var (
	ErrAppRunning    = errors.New(`application is already running`)
	ErrAppNotRunning = errors.New(`application is not running `)
)
