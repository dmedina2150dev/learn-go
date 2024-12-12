## Usos de los controles de flujos

* Operadores booleanos
* Declaración de IF - ELSE
* Declaración de Switch
* Bucle For
* Uso de funciones
* Proyecto de la sección


### Operadores booleanos

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


### Declaración de IF - ELSE

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
