# LEARN GO

### A) Trabajando con Paquetes externos de GO

<br>
Los paquetes externos de GO, son paquetes creados por la comunidad de desarrolladores los cuales no son mantenidos por el equipo de GO.

<br>

Para utilizarlos debemos tenemos que inicializar un manejador de Modulos para nuestra aplicación.

## Como inicializar el manejador de paquetes

```cmd
go mod init [nombre-del-manejador-de-modulos]
```

Esto nos creará un archivo con el nombre **go.mod** este archivo contendra

1. EL nombre del manejador de modulo que asignamos al momento de ejecutar el comando.
2. Una version de **GO**

```go
module appone

go 1.23.0
```

Este archivo es utilizado para definir y gestionar los modulos y las dependencias del proyecto.

> **NOTA: Cuando descargemos y utilicemos archivos externos se nos creara un archivo que sera \_\_\***go.sum**\*\_\_, este archivo se utiliza para registrar la suma de verificacionesde los modulos y las dependencias de nuestro proyecto**

### A-1) Usando el paquete (quote)

Es un paquete externo o de terceros, que proporciona una serie de citas famosas como el "Hola Mundo"

```cmd
go get rsc.io/quote
```

## Para descargar un paquete en GO

```cmd
go get [ruta-repositorio-paquete]
```

**NOTA: El paquete podrias estar en Github o en otro repositorio o URL**

Esto descarga el paquete y lo agrega en nuestro manejador de paquetes, que quedaria de la siguiente forma.

```go
module appone

go 1.23.0

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.3.0 // indirect
)
```

## Importar paquetes

Para un solo paquete

```go
import "fmt"
```

Para importar varios paquetes

```go
import (
  "fmt"
)
```

## Comandos para ejecutar programa

> Descarga las dependencias del proyecto

```go
go mod tidy

```

> Ejecuta la aplicación en local

```go
go run [name-main-file].go

```

> Compila la aplicación en un ejecutable

```go
go build [name-main-file].go

```

## Comandos para ejecutar los tests

### Para ejecutar todos los tests del proyecto

```go
go test ./...

```

### Para ejecutar todos los tests del proyecto con verbose

Verbose da más detalle interno de cada test

```go
go test ./... -v

```

### Para ver el coverage del proyecto 
```go
go test -cover ./...

```

### Coverage detallado con reporte HTML

Un archivo coverage.out con los datos de coverage
Un archivo coverage.html que puedes abrir en el navegador para ver el coverage línea por línea

```go
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Coverage con porcentaje total

El flag -func muestra el coverage por función y el total al final.

```go
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

### Para un paquete específico

```go
go test -cover ./ruta_relativa_del_paquete
```

### Coverage con verbose para más detalles

```go
go test -v -cover ./...
```

###  Coverage con umbral mínimo

Puedes crear un script que falle si el coverage está por debajo de cierto porcentaje:

```go
go test -coverprofile=coverage.out ./...
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
if (( $(echo "$COVERAGE < 80" | bc -l) )); then
    echo "Coverage is below 80%: $COVERAGE%"
    exit 1
fi
```

