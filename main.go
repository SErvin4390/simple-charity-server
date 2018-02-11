package main

import (
	"fmt"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_chi"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	fmt.Println("\nStarting Simple Charity Server")
	rand.Seed(time.Now().UTC().UnixNano())

	router := SetupApp()

	fmt.Printf("\n\t Listening on port: %s\n", Config.Port)
	http.ListenAndServe(Config.Port, router)
}

//SetupApp sets up the Chi Router and basic configuration for the application
func SetupApp() *chi.Mux {
	Config = ConfigSetup()

	r := chi.NewRouter()
	limiter := tollbooth.NewLimiter(100.0, nil)
	render.SetContentType(render.ContentTypeJSON)
	middleware.Timeout(120 * time.Second)

	r.Use(middleware.StripSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(tollbooth_chi.LimitHandler(limiter))

	//set all routes here
	r.Get("/", NotImplementedRoute)

	return r
}

//NotImplementedRoute is a simple route that simply returns a 501
func NotImplementedRoute(w http.ResponseWriter, r *http.Request) {
	Send(w, 501, "Route Not Implemented")
	return
}
