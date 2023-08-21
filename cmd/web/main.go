package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const portNumber = "8080"

var app config.AppConfig
var sesion *scs.SessiionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	// session
	sesion = scs.New()
	sesion.Lifetime = 24 * time.Hour
	sesion.Cookie.Persist = true
	sesion.Cookie.Secure = app.InProduction

	app.Session = sesion

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
