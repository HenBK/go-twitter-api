package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/henbk/go-twitter-api/middlewares"
	"github.com/henbk/go-twitter-api/routers"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	// TODO: create helper function that receives as parameters a vector of middleware functions and a handler function and returns
	// the handler function with the middleware functions applied, just to avoid middleware functions nesting too deep

	router.HandleFunc("/register", middlewares.DatabaseConnectionCheck(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.DatabaseConnectionCheck(routers.Login)).Methods("POST")
	router.HandleFunc("/profile",
		middlewares.DatabaseConnectionCheck(
			middlewares.ValidateJsonWebToken(routers.Profile),
		),
	).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
