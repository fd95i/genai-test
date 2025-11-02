package main

import (
	"log"

	"genai-test/infrastructure/config"
)

func main() {
	// Construir la aplicación mediante el factory
	app, err := config.BuildApplication()
	if err != nil {
		log.Fatalf("Error al construir la aplicación: %v", err)
	}

	// Ejecutar la aplicación
	if err := app.Run(); err != nil {
		log.Fatalf("Error al ejecutar la aplicación: %v", err)
	}
}
