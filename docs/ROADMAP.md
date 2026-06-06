🧭 Ruta de aprendizaje Go (Golang) para Backend Developers

🧱 Módulo 1: Fundamentos del lenguaje Go

Objetivo: Familiarizarte con la sintaxis y estructuras básicas de Go.

Contenido: 1. Instalación y primer programa (Hello World) 2. Variables y constantes 3. Tipos de datos (string, int, bool, float, etc.) 4. Funciones y parámetros 5. Estructuras de control (if, switch, for, range) 6. Arrays, slices y maps 7. Structs y métodos 8. Interfaces y composición 9. Manejo básico de errores

👉 Ejercicio práctico:
• Crear un programa que gestione una lista de reservas en memoria (crear, listar, eliminar).
• Usa structs, slices, map, funcs y interfaces.

🎯 Módulo 2: Buenas prácticas y entorno Go

Objetivo: Trabajar como un Gopher real, entendiendo la organización del código y herramientas esenciales.

Contenido: 1. Estilo idiomático (gofmt, go vet, go lint) 2. Organización de paquetes y módulos (go mod init, imports) 3. Estructura estándar de un proyecto Go 4. Introducción a testing (testing package) 5. Herramientas útiles (testify, golangci-lint, air para hot reload)

👉 Ejercicio práctico:
• Refactoriza el proyecto anterior en una estructura modular.
• Escribe un par de tests unitarios simples.

⸻

🌐 Módulo 3: Backend puro con Go

Objetivo: Crear tu primera API RESTful con Go desde cero.

Contenido: 1. Introducción a net/http 2. Servidor básico, rutas y handlers 3. Manejo de request/response, encoding JSON 4. Validación básica y middlewares manuales 5. Introducción a Gin (router más expresivo y potente)

👉 Ejercicio práctico:
• API REST con CRUD de reservas:
• GET /reservas
• POST /reservas
• DELETE /reservas/:id
• Usa Gin + middlewares personalizados + JSON binding

⸻

🛢️ Módulo 4: Bases de datos y persistencia

Objetivo: Conectar tu backend a una base de datos real (PostgreSQL o MySQL).

Contenido: 1. database/sql + driver nativo (pgx) 2. Consultas básicas con Query y Exec 3. Uso de sqlx para queries más limpias 4. Alternativas: GORM, sqlc 5. Migraciones de schema con golang-migrate

👉 Ejercicio práctico:
• Persistir las reservas en una base de datos PostgreSQL.
• Escribir funciones para Insert, GetAll, DeleteByID.
• Añadir migraciones automáticas con golang-migrate.

⸻

⚙️ Módulo 5: Concurrencia y arquitectura

Objetivo: Aprender herramientas avanzadas para un backend robusto.

Contenido: 1. Goroutines y canales (go, chan) 2. context.Context en handlers y DB queries 3. Logging profesional con zerolog o logrus 4. Errores personalizados y wrapping 5. Arquitectura limpia: separación por capas (cmd, internal, pkg)

👉 Ejercicio práctico:
• Aplicar arquitectura limpia a tu API de reservas.
• Añadir logs estructurados y contexto a cada request.

⸻

🚀 Módulo 6: Preparación para producción

Objetivo: Dejar tu backend listo para producción.

Contenido: 1. Configuración por variables de entorno (os.Getenv, godotenv) 2. Crear imagen Docker de tu API 3. Despliegue en Railway o Fly.io 4. Seguridad básica (headers, CORS, validaciones) 5. Buenas prácticas de rendimiento

👉 Ejercicio final:
• Dockeriza tu API.
• Sube el proyecto a Railway.
• Expón tu documentación con Swagger o README bien estructurado.
