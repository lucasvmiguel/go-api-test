package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/lucasvmiguel/go-api-test/db"
	"github.com/lucasvmiguel/go-api-test/internal/post"
	"github.com/lucasvmiguel/go-api-test/internal/todo"
	"github.com/lucasvmiguel/go-api-test/pkg/cmd"
	"github.com/lucasvmiguel/go-api-test/pkg/ping"
)

type config struct {
	Port int `env:"PORT" envDefault:"8080"`
}

func main() {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		cmd.ExitWithError("failed to read config", err)
	}

	dbClient := db.NewClient()
	err = dbClient.Prisma.Connect()
	if err != nil {
		cmd.ExitWithError("failed to connect to the DB", err)
	}

	defer func() {
		err := dbClient.Prisma.Disconnect()
		if err != nil {
			cmd.ExitWithError("failed to disconnect from the DB", err)
		}
	}()

	router := chi.NewRouter()

	// middlewaress
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request
	router.Use(middleware.Timeout(60 * time.Second))

	// ping handler
	router.Handle("/ping", &ping.Handler{})

	// todo handler
	todoHandler, err := todo.NewHandler(&http.Client{}, "https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		cmd.ExitWithError("todo handler had an error", err)
	}
	router.Get("/todos", todoHandler.ServeHTTP)

	// post handler [POST]
	postHandlerPost, err := post.NewHandlerPost(dbClient)
	if err != nil {
		cmd.ExitWithError("post handler post had an error", err)
	}
	router.Post("/posts", postHandlerPost.ServeHTTP)

	// post handler [GET]
	postHandlerGet, err := post.NewHandlerGet(dbClient)
	if err != nil {
		cmd.ExitWithError("post handler get had an error", err)
	}
	router.Get("/posts", postHandlerGet.ServeHTTP)

	log.Printf("listening on port %d", cfg.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router))
}
