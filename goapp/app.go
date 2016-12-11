package goapp

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/burntsushi/toml"
	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

// App is main web application
type App struct {
	*httprouter.Router              // httprouter
	handlers           http.Handler // modified net/http handler (for middlewares)
	config             Config       // configuration object
	db                 *MongoClient // database client (pointer used to avoid session copy)
}

// NewApp creates new web application
func NewApp(configFile string) *App {

	// initialize empty App struct
	app := App{
		Router: httprouter.New(),
	}

	// load config file if exists
	if configFile != "" {
		if _, err := toml.DecodeFile(configFile, &app.config); err != nil {
			log.Panic(err)
		}
	}

	// establish db connection
	app.db = NewMongoClient(app.config.Database.URI, "goapp")

	// add handlers
	app.GET("/", app.index)

	// add middlewares
	h := handlers.LoggingHandler(os.Stdout, app)
	h = handlers.ProxyHeaders(h)
	h = handlers.CompressHandler(h)
	h = handlers.RecoveryHandler()(h)
	app.handlers = h

	return &app
}

// Handle handler errors
func (a *App) handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// Handle bad request
func (a *App) handleBadRequest(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusBadRequest)
}

// Run runs this web application
func (a *App) Run() {
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", a.config.Server.Port), a.handlers),
	)
}
