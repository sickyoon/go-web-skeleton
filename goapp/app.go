package goapp

import (
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
}

// NewApp creates new web application
func NewApp(configFile string) *App {

	// initialize empty App struct
	app := App{
		Router: httprouter.New(),
	}

	// load config file if exists
	if configFile != "" {
		if _, err := toml.DecodeFile(configFile, app.config); err != nil {
			log.Panic(err)
		}
	}

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

// Run runs this web application
// TODO: do escape analysis to see if pointer receiver makes sense
//       value receiver saves memory in certain cases
func (a App) Run() {
	log.Fatal(http.ListenAndServe(":8000", a.handlers))
}
