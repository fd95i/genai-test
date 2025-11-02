package application

import (
	"genai-test/infrastructure/router"
)

// App es la implementación concreta de Application
type App struct {
	router *router.Router
}

// NewApp crea una nueva instancia de App
func NewApp(r *router.Router) *App {
	return &App{
		router: r,
	}
}

// Start inicializa la aplicación
func (a *App) Start() error {
	// Aquí se pueden inicializar servicios, conexiones a BD, etc.
	return nil
}

// Run inicia el servidor HTTP
func (a *App) Run() error {
	return a.router.GetEngine().Run()
}
