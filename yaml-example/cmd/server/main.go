// GENERATED CODE, DO NOT EDIT!
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cheikhshift/samb-examples/yaml-example/pkg/api"
	"github.com/cheikhshift/samb-examples/yaml-example/pkg/hooks"
)

func main() {

	hooks.Start()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt)

	http.HandleFunc("/", api.Handler)

	h := &http.Server{Addr: host + ":" + port}

	go func() {
		<-stop
		cleanUp(h)
	}()

	err := h.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}

func cleanUp(h *http.Server) {
	log.Println("\nShutting down the server...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	h.Shutdown(ctx)

	hooks.Stop()
	log.Println("Server gracefully stopped")

}
