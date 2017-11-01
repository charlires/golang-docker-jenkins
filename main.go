package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "db:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_ = client.Set("user", 0, 0).Err()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		val, _ := client.Get("user").Result()
		w.WriteHeader(http.StatusOK)
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "Hello, World!\nHostname: %s\nUser Number:%s\n", hostname, val)
	})

	r.HandleFunc("/redis", func(w http.ResponseWriter, r *http.Request) {
		val, err := client.Get("user").Result()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "No redis :'(\nError:", err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, visitor :%s\n", val)
		_ = client.Incr("user").Err()
	})

	fmt.Println("running server")
	if err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r)); err != nil {
		fmt.Printf("unable to start server. %+v", err)
		return
	}
}
