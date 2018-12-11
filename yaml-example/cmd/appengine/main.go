// GENERATED CODE, DO NOT EDIT!
package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/cheikhshift/samb-examples/yaml-example/pkg/api"
	"github.com/cheikhshift/samb-examples/yaml-example/pkg/hooks"
	"google.golang.org/appengine"
)

func main() {

	hooks.Start()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt)

	http.HandleFunc("/", api.Handler)

	go func() {
		<-stop
		cleanUp()
	}()

	// Launch app on appengine
	appengine.Main()
}

func cleanUp() {

	hooks.Stop()
	log.Println("App gracefully stopped")

}
