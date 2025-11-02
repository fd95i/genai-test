package config

import (
	"genai-test/application"
	appInfra "genai-test/infrastructure/application"
	"genai-test/infrastructure/router"
)

// BuildApplication construye todas las capas de la aplicación y devuelve
// una instancia de Application correctamente inicializada
func BuildApplication() (application.Application, error) {
	// Construir Router
	r := router.NewRouter()

	// Construir Application
	app := appInfra.NewApp(r)

	// Inicializar la aplicación
	if err := app.Start(); err != nil {
		return nil, err
	}

	return app, nil
}
