# Golang-Concept-Core

# Understanding Arrays in Go

Alright! imagine you have a row of watch boxes, each box holding a watch sonata.this row is like an array in golang.An array is a collection of items,all of the same type,arranged in a specific order.

## What is an Array in Go?------------

An **array** in go is a **fixed size** sequence of elements of the **same type**.once you define the size of an array, it **can not change**.

### How to use Arrays with examples------------

1. ## Declaring and initializing an Array

you can declare an array by specifying its size and type, and optionally initialize it with values.
```go
var numbers [3]int = [3]int{1,2,3}

Or, using shorthand:
```go
numbers := [3]int{1,2,3}
```
Here, numbers is an array of three integers: 1,2, and 3

2. ## Accessing Array Elements

you can access element of an array using their index, starting from 0.
```go
fmt.Println(numbers[0])  //outputs: 1
```

3. ## Modifying Array Elements
you can change the value of an **array element** by assigning a new value to its **index**.
```go
numbers[1] = 10
fmt.Println(numbers) //outputs: [1,10,3]
```

### Types Of Arrays in Go:------------

in Go, arrays can be two types.

1. ## One-dimensinal Arrays
A simple list of elements.
```go
var letters [3]string = [3]string{"A", "B", "C"}
```

2. ## Multi-dimensinal Arrays
Arrays containing other arrays, like a table of rows and columns.
```go
var matrix [2][2]int = [2][2]int{{1,2}, {3,4}}
```

### Declaring Arrays in Go
there are several ways to declare arrays in Go:

1. ## Using a var keyword with ecplicit length
Specify the length and type explicitly.
```go
var fruits [2]string
fruits[0] = "Apple"
fruits[1] = "Banana"
```

2. ## Using shorthands decleration with initialization
let Go infer the length by initializing with values.
```go
colors := [...]string{"Apple", "Banana", "Kivi"}
```

Here, Go determines the array length based on the number of elements privided.

3. ## Partially Initialized Arrays
you can initialize some elements, and the rest will have zero values.
```go
var ages [3]int = [3]int{5,10}
fmt.Println(ages)  //outputs: [5,10,0]
```

The third element is 0 becouse it's the zero value for integers.

4. ## Multi-dimensinal Arrays
Declare arrays within arrays.
```go
var grid [2][3]int
```
this creates a 2*3 grid(2 rpws, 3 columns) of integers.