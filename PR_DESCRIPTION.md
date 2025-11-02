# Feature: Hexagonal Architecture Scaffolding

## 📋 Resumen

Implementación inicial del scaffolding del proyecto utilizando arquitectura hexagonal (ports & adapters) con Go, Gin como router, y configuración para testing con Ginkgo v2 y Gomega.

## 🏗️ Arquitectura Implementada

El proyecto sigue una arquitectura hexagonal con las siguientes capas:

### 1. **Application Layer** (`application/`)
- Contiene la lógica de negocio, casos de uso y puertos (interfaces)
- Define la interfaz `Application` con métodos `Start()` y `Run()`

### 2. **Domain Layer** (`domain/`)
- Preparado para contener las entidades de dominio
- Actualmente vacío, listo para futuras implementaciones

### 3. **Infrastructure Layer** (`infrastructure/`)
- **`router/`**: Adaptador HTTP utilizando Gin
  - Implementa el struct `Router` que encapsula Gin
  - Configura ruta `/health` GET que retorna `{"status": "UP"}`
- **`application/`**: Implementación concreta de la interfaz `Application`
  - Struct `App` que implementa `application.Application`
- **`config/`**: Factory pattern para construcción de dependencias
  - `factory.go` con función `BuildApplication()` que construye toda la estructura

## 📁 Estructura de Archivos

```
genai-test/
├── application/
│   └── application.go          # Interfaz Application
├── domain/                      # Entidades de dominio (preparado)
├── infrastructure/
│   ├── application/
│   │   └── app.go              # Implementación de Application
│   ├── config/
│   │   └── factory.go          # Factory para construcción de dependencias
│   └── router/
│       └── router.go           # Router con Gin
└── main.go                      # Punto de entrada
```

## 🔧 Componentes Implementados

### Interfaz Application
```go
type Application interface {
    Start() error
    Run() error
}
```

### Router
- Struct `Router` que encapsula `*gin.Engine`
- Método `NewRouter()` que crea e inicializa el router
- Ruta `/health` configurada para health checks

### Factory Pattern
- `BuildApplication()` construye todas las capas:
  1. Crea el Router mediante `router.NewRouter()`
  2. Crea la Application mediante `appInfra.NewApp(r)`
  3. Inicializa la aplicación con `app.Start()`
  4. Retorna la instancia lista para ejecutar

### Main
- Invoca `Application.Run()` mediante la instancia obtenida del factory

## 🧪 Testing Setup

- Dependencias instaladas: Ginkgo v2 y Gomega
- Configuración de tests pendiente (siguiente paso)

## ✅ Verificaciones

- ✅ Proyecto compila sin errores
- ✅ Servidor HTTP inicia correctamente en puerto 8080
- ✅ Endpoint `/health` funciona y retorna `{"status": "UP"}`

## 📝 Prompts Utilizados

### Prompt Inicial
```
Quiero que construyamos un proyecto en Go utilizando Gin como router, Ginkgo v2 y Gomega para testing. Vamos a comenzar con el scaffolding, te parece?
```

### Prompt de Arquitectura
```
me gustaría utilizar arquitectura hexagonal, con las siguientes capas:

application: lógica de negocio, casos de uso, puertos (interfaces de E/S)
domain: entidades de dominio principalmente
infrastructure: adaptadores, rutas, configuraciones externas

el archivo main debería invocar un método Run definido en una interfaz Application, la cual se inicializará mediante un método start y construirá la estructura de la siguiente manera:

existirá un archivo factory.go ubicado en infrastructure/config que se encargará de construír todas las capas de la aplicación y devolverá un Application correctamente inicializado y listo para ejecutar Run.

Comencemos creando los directorios correspondientes, el struct Application con sus métodos Start() y Run() y un struct Router que, utilizando gin, contenga una ruta /health GET que retorne un json con una key "status" y un valor "UP". El archivo factory.go debería inicializar el Application con su respectivo Router mediante el método Start() (para Application) y el método NewRouter() (para Router)
```

## 🚀 Próximos Pasos

- [ ] Configurar tests con Ginkgo v2 y Gomega
- [ ] Implementar casos de uso en la capa Application
- [ ] Agregar entidades de dominio
- [ ] Implementar más endpoints y handlers

## 📦 Dependencias

- `github.com/gin-gonic/gin v1.11.0`
- `github.com/onsi/ginkgo/v2 v2.27.2`
- `github.com/onsi/gomega v1.38.2`

