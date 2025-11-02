## 📊 Cambios Realizados

Se han agregado diagramas de flujo completos para documentar la arquitectura hexagonal del proyecto utilizando Mermaid.

### ✨ Nuevos Diagramas Agregados:

1. **Diagrama General de Arquitectura Hexagonal**
   - Muestra la relación entre todas las capas (Domain, Application, Infrastructure)
   - Ilustra las dependencias y flujos entre componentes
   - Color coding diferenciado por capa:
     - 🔵 Azul claro: Infrastructure Layer
     - 🟡 Amarillo claro: Application Layer  
     - 🩷 Rosa claro: Domain Layer

2. **Flujo de la Capa Domain**
   - Proceso: Entities → Rules → Validation → Events
   - Representa la lógica de negocio pura sin dependencias externas

3. **Flujo de la Capa Application**
   - Define interfaces y contratos (ports)
   - Casos de uso y orquestación de lógica de negocio

4. **Flujos de la Capa Infrastructure** (3 sub-diagramas detallados):
   
   **a. Factory Flow:**
   - Proceso de construcción paso a paso de la aplicación
   - Construcción de Router → App → Inicialización
   - Manejo de errores en cada etapa
   
   **b. Router Flow:**
   - Inicialización del router Gin
   - Configuración de rutas
   - Manejo de requests HTTP y routing
   
   **c. Application Implementation Flow:**
   - Ciclo de vida completo: NewApp → Start → Run
   - Inicialización de servicios
   - Puesta en marcha del servidor HTTP

5. **Flujo de Ejecución Completo (Sequence Diagram)**
   - Diagrama de secuencia que muestra todas las interacciones
   - Desde `main.go` hasta la respuesta HTTP al cliente
   - Incluye todos los participantes: Main, Factory, Router, App, Gin Engine, HTTP Client

### 📝 Mejoras Implementadas:

- ✅ Diagramas en formato **Mermaid** (compatible con GitHub, GitLab y visores de Markdown)
- ✅ Documentación visual completa de cada capa de la arquitectura
- ✅ Índice del README actualizado con las nuevas secciones
- ✅ Flujos detallados que facilitan el entendimiento de la arquitectura
- ✅ Separación clara de responsabilidades visualizada

### 📈 Estadísticas:

- **Líneas agregadas**: +184
- **Líneas modificadas**: -18
- **Total**: +166 líneas de documentación visual

### 🎯 Impacto:

Los diagramas proporcionan:
- **Mejor comprensión** de la arquitectura hexagonal para nuevos desarrolladores
- **Documentación visual** que complementa la documentación textual
- **Referencia rápida** para entender el flujo de ejecución
- **Facilita el onboarding** de nuevos miembros del equipo

