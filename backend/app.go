package main

import (
	"net/http"
	//Router Chi
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	//Module routes
	"github.com/walteralv/enronMailApp/services"
)


func main() {

	app := chi.NewRouter()

	//Middlewares
	app.Use(middleware.Logger)
	app.Use(middleware.AllowContentType("application/json", "text/xml"))

	app.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: false,
	}))

	// Adding mores routes to app
	app.Get("/emails/search", services.SearchEmail)

	http.ListenAndServe(":8000", app)
}
