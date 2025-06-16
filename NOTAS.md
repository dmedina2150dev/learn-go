# Notas que para mi son relevantes en GO

### Declaraciones publicas

En **GO** cuando queremos que una `func` pueda ser accedida desde fuera del File donde es declarado debemos escribir su nombre comenzando con mayuscula.

```go
package prueba

func ProbarConn() {}

func connect() {}

```

```go
package main

import "[rutapackage.prueba]"

func main() {

    // Acá podremos ver invocar esta metodo
    prueba.ProbarConn()

    // Acá no podremos acceder a esté
    prueba.connect()

}

```

Pasa algo similar con las variables, más adelante se explica.

### Lo que pensaba que era Magia

En **Go** como en otros lenguajes podemos definir folders y dentro de estas crear archivos que subdividen nuestra logica de negocio o las funcionalidades que estamos desarrollando. En GO se utiliza el `package [name-package]` internamente dentro del folder no sera necesario importar cada uno de los archivos que usemos compartidos.

> folder: [variables/enteros.go]

```go
package variables

import "fmt"

func MuestroEnteros() {
	intde32 := int32(10)

	fmt.Println("intde32 = ", intde32)
}
```

> folder: [variables/resto.go]

```go
package variables

import "fmt"

func RestoVariables() {
	MuestroEnteros()
	fmt.Println("No se tiene que importar el archivo donde está la función")
}
```

### Scope de las variables

Con las variables tambien podemos hacer algo similar a las `func` se pueden compartir internamente en el package o permitir que se usen en otros package siempre que se cumplan las condiciones:

- Para usar entre package

  - Debe estar definida fuera del scope de una `func`
  - Su nombre debe empezar con mayuscula

- Para usar solo en el package definida
  - Debe estar definida fuera del scope de una `func`
  - Su nombre debe empezar con minusculas

> folder: [variables/enteros.go]

```go
package variables

import "fmt"

var num1 float32
var Num2 float64

func MuestroEnteros() {
	intde32 := int32(10)
    Num3 := 20

	fmt.Println("intde32 = ", intde32)
}
```

> folder: [variables/resto.go]

```go
package variables

import "fmt"

func RestoVariables() {
	MuestroEnteros()
    fmt.Println("Acá podemos mostrar: ", num1, Num2)

	// fmt.Println("No se puede mostrar", Num3)
}
```

### Guía de estilos para programar en GO

[Github](https://github.com/uber-go/guide/blob/master/style.md#test-tables)

### Diferencias entre un `slice` y un `array`

Un `slice` es una abstracción de un array que permite trabajar con una colección de elementos de manera más flexible. A diferencia de un array, un slice no tiene un tamaño fijo y puede crecer o decrecer dinámicamente. Además, los slices son más fáciles de manipular y proporcionan una interfaz más rica para trabajar con colecciones de datos.
Un array es una colección de elementos del mismo tipo con un tamaño fijo, mientras que un slice es una vista dinámica sobre un array subyacente. Los slices permiten trabajar con colecciones de datos de manera más flexible y eficiente, ya que pueden crecer o decrecer según sea necesario.

## make

La función `make` en Go se utiliza para crear slices, maps y channels. A diferencia de `new`, que asigna memoria, `make` inicializa el objeto y lo prepara para su uso. Y podemos indicarle el tamaño inicial de estos tipos de datos y también su capacidad.

### Uso de `make` con slices

```go
slice := make([]int, 0)
```

### Uso de `make` con maps

```go
m := make(map[string]int)
```

### Uso de `make` con channels

```go
ch := make(chan int)
```

### Conocer la capacidad de un slice

Para conocer la capacidad de un slice, se puede utilizar la función `cap(slice)`, que devuelve la capacidad del slice, es decir, el número de elementos que puede contener sin necesidad de redimensionarse.

```go
package main
import "fmt"
func main() {
    slice := make([]int, 0, 5) // Crea un slice con capacidad para 5 elementos
    fmt.Println("Capacidad del slice:", cap(slice)) // Imprime: Capacidad del slice: 5
}
```
