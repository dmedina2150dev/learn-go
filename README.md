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

Esto nos creará un archivo con el nombre __go.mod__ este archivo contendra

1. EL nombre del manejador de modulo que asignamos al momento de ejecutar el comando.
2. Una version de __GO__

```go
module appone

go 1.23.0
```

Este archivo es utilizado para definir y gestionar los modulos y las dependencias del proyecto.

> **NOTA: Cuando descargemos y utilicemos archivos externos se nos creara un archivo que sera __***go.sum***__, este archivo se utiliza para registrar la suma de verificacionesde los modulos y las dependencias de nuestro proyecto**


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

> Levanta la aplicación en local

```go
go run [name-main-file].go

```

> Compila la aplicación en un ejecutable

```go
go build [name-main-file].go

```