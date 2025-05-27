# Notas que para mi son relevantes en GO

### Declaraciones publicas

En __GO__ cuando queremos que una `func` pueda ser accedida desde fuera del File donde es declarado debemos escribir su nombre comenzando con mayuscula.

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

En __Go__ como en otros lenguajes podemos definir folders y dentro de estas crear archivos que subdividen nuestra logica de negocio o las funcionalidades que estamos desarrollando. En GO se utiliza el `package [name-package]` internamente dentro del folder no sera necesario importar cada uno de los archivos que usemos compartidos.

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
    * Debe estar definida fuera del scope de una `func`
    * Su nombre debe empezar con mayuscula

- Para usar solo en el package definida
    * Debe estar definida fuera del scope de una `func`
    * Su nombre debe empezar con minusculas

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
