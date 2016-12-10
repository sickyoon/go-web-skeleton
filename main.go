package main

import (
	"flag"
	"runtime"

	"github.com/sickyoon/go-web-skeleton/goapp"
)

var config = flag.String("config", "", "configuration file")

func main() {
	// parse cmd args
	flag.Parse()

	// set proc num
	runtime.GOMAXPROCS(runtime.NumCPU())

	// create new web application
	app := goapp.NewApp(*config)

	// run application
	// TODO: graceful shutdown
	app.Run()
}
