# Usos de los controles de flujos

* Operadores booleanos
* Declaración de IF - ELSE
* Declaración de Switch
* Bucle For
* Uso de funciones
* Proyecto de la sección

---------------
## Operadores booleanos

Los operadores relacionales y lógicos son utilizados en conjunto en las expresiones lógicas de __GO__ para evaluar condiciones complejas y producir un resultado booleano (verdadero o falso).

#### __Operadores de comparación__:

Se usan para comparar dos valores y devolver un valor booleano (true o false) según el resultado de la comparación. Los operadores de comparación incluyen:

> Igualdad
```
(==)
```

> Desigualdad
```
(!=)
```

> Mayor que
```
(>)
```

> Menor que
```
(<)
```

> Mayor o igual que
```
(>=)
```

> Menor o igual que
```
(<=)
```


#### __Operadores lógicos__:

Los operadores lógicos en __Go__ son utilizados para evaluar expresiones lógicas y producir un resultado booleano (veradero o falso). En __Go__, existen 3 (tres) operadores lógicos

----------------
> Operador AND lógico (&&) 

Evalúa dos expresiones booleanas y devuelve verdader (true) si ambas expresiones son verdaderas, y devuelve falso (false) si algúna de las expresiones es falsa

```
(&&)
```

-----------------
> Operador OR lógico (||)

Evalúa dos expresiones booleanas y devuelve verdader (true) si al menos una de las expresiones es verdadera, y devuelve falso (false) si ambas expresiones son falas.

```
(||)
```

-----------------
> Operador NOT lógico (!)

Niega el valor booleano de una expresión, es decir, si una expresión es verdadera, la niega y devuelve falso, y si una expresión es falsa, la niega y devuelve verdadero.

```
(!)
```

---------------
## Declaración de IF - ELSE

Se utilizan para ejecutar un bloque de código si se cumple una condición booleana, y si no se cumple no se ejecuto o se ejecuta la segunda condición del bloque de código.

__NOTA:__ La condición que se evalúa en un IF en __Go__ se debe colocar sin parentesís a diferencia de otros lenguajes de programación.


```Go
if condicion {
    // Si se cumple la condición se ejecuta el bloque de código
} else {
    // en caso de que no se cumpla la condición
}
```

De igual forma podemos evalúar más condiciones con ELSE IF

```Go
if condicion {
    // Si se cumple la condición se ejecuta el bloque de código
} else if codicion2 {
    // en caso de que no se cumpla la primera condición
} else {
    // en caso de que no se cumpla ningúna de las condiciones
}
```

__NOTA__: En __Go__ es posible definir o declarar e iniciarlizar las variables dentro de una condición y esta variable solo estará disponible en el scope del IF.

```Go
if tiempo := time.Now(); tiempo.Hour() < 12 {
    fmt.Println("Buenos días - Es de mañana")
} else if tiempo.Hour() < 17 {
    fmt.Println("Buenas tardes - Es de tarde")
} else {
    fmt.Println("Buenas noches - Es de noche")
}
```

---------------
## Declaración de Switch

Se utiliza colocando la palabra reservada __switch__ está nos permitira evalúar un valor en base a diferentes casos __case__ definidos. Cuando el valor coincida con algúno de los casos se ejcutara el bloque de código contenido dentro del caso.

Al igual que todos los lenguajes contiene una __default__ donde se ejecutara ese bloque, en el momento en que no se cumpla o no coincida con los __case__ definidos

__NOTA__: Destacar que aunque exite la expresión __break__ para detener la ejecución en ese punto ya no es necesario.

```Go
switch valor {
case useCase:
    // bloque de código a ejecutar si cumple
default:
    // Si no hay caso que coicida
}
```

__NOTA__: En __Go__ es posible definir o declarar e iniciarlizar las variables dentro de una condición y esta variable solo estará disponible en el scope del switch

```Go
switch system := runtime.GOOS; system {
case "windows":
    fmt.Println("Windows.")
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    fmt.Printf("%s.\n", system)
}
```

## Bucle For

En __Go__ el bucle FOR es la unica estructura de control repetitiva disponible, pero esta puede ser usada de diferentes maneras, para conseguir diferentes funcionalidades. Como crear un bucle infinito, un bucle con una condición o el tipico Bucle For, como tambien iterar colecciones de datos.


> Declarar un bucle infinito

```Go
for {

}
```

> Declarar bucle con una condición

```Go
for condición {

}
```

Ejemplos: 

```Go
var i int

for i <= 10 {
    fmt.Println(i)
    i++
}

for j := 0; j <= 10; j++ {
    fmt.Println(j)
}
```

### Break
Es una palabra clave del lenguaje __Go__, se utiliza para salir de un bucle antes de que la condición de finalización se haya alcanzado

```Go
for j := 0; j <= 10; j++ {
    fmt.Println(j)
    if j == 5 {
        break
    }
}
```

### Continue
Esta palabra clave, se utiliza para saltar a la siguiente iteración de un bucle sin ejecutar el coódigo que esta despues de esta sentencia o palabra clave.

```Go
for j := 0; j <= 10; j++ {
    if j == 5 {
        continue
    }
    fmt.Println(j)
}
```
