package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"

	"psubbin/internal/assets"
)

func main() {
	port := ":3000"

	flag.StringVar(&port, "port", port, "port to listen on")
	flag.Parse()

	router := chi.NewRouter()
	// list := domain.NewMessage("Hello, world!", "home")
	// list.Add("Bake a cake")
	// list.Add("Feed the cat")
	// list.Add("Take out the trash")

	// home.Mount(router, home.NewHandler(home.NewService(list)))
	// todos.Mount(router, todos.NewHandler(todos.NewService(list)))
	assets.Mount(router)

	server := &http.Server{
		Addr:    port,
		Handler: http.TimeoutHandler(router, 30*time.Second, "request timed out"),
	}

	// Display the localhost address and port
	fmt.Printf("Listening on http://localhost%s\n", port)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
