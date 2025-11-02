package router

import (
	"github.com/gin-gonic/gin"
)

// Router encapsula el router de Gin
type Router struct {
	engine *gin.Engine
}

// NewRouter crea una nueva instancia del Router con Gin
func NewRouter() *Router {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	router := &Router{
		engine: engine,
	}

	router.setupRoutes()

	return router
}

// setupRoutes configura las rutas de la aplicación
func (r *Router) setupRoutes() {
	r.engine.GET("/health", r.healthHandler)
}

// healthHandler maneja la ruta /health
func (r *Router) healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}

// GetEngine retorna el engine de Gin para uso interno
func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}
