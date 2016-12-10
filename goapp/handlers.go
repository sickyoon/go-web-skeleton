package goapp

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *App) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hello")
}
