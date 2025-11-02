package application

// Application define la interfaz principal de la aplicación
type Application interface {
	Start() error
	Run() error
}
