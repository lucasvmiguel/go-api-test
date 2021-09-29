package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/lucasvmiguel/go-api-test/db"
	"github.com/lucasvmiguel/go-api-test/internal/post"
	"github.com/lucasvmiguel/go-api-test/internal/todo"
	"github.com/lucasvmiguel/go-api-test/pkg/cmd"
	"github.com/lucasvmiguel/go-api-test/pkg/ping"
)

func main() {
	dbClient := db.NewClient()
	err := dbClient.Prisma.Connect()
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
	router.Get("/todos", todoHandler.Handle)

	// post handler
	postHandler, err := post.NewHandler(dbClient)
	if err != nil {
		cmd.ExitWithError("post handler had an error", err)
	}
	router.Route("/posts", func(r chi.Router) {
		r.Post("/", postHandler.HandlePost)
		r.Get("/", postHandler.HandleGet)
	})

	http.ListenAndServe(":8080", router)
}
