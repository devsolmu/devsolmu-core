package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devsolmu/devsolmu-core/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instance
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (app *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, error := gorm.Open(config.DB.Dialect, dbURI)
	if error != nil {
		log.Fatal("Could not connect database")
	}

	app.DB = DBMigrate(db)
	app.Router = mux.NewRouter()
	app.setRouters()
}

// setRouters sets the all required routers
func (app *App) setRouters() {
	// Routing for handling the projects
	app.Get("/projects", app.GetAllProjects)
	app.Get("/projects", app.GetProject)
	app.Post("/projects", app.CreateProject)
	app.Put("/projects", app.UpdateProject)
	app.Delete("/projects", app.DeleteProject)

	// Others...
}

// Get wraps the router for GET method
func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (app *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (app *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}
