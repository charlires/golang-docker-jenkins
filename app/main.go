package main

import (
	"fmt"
	"net/http"
	"os"

	"webapp/app/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Main)
	r.HandleFunc("/healthz", handler.Healthz)

	fmt.Println("running server")
	if err := http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, r)); err != nil {
		fmt.Printf("unable to start server. %+v", err)
		return
	}
}
