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
```

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


# Understanding Slice in Go

imagine you have a magical watch box that can change its size whenever you want. if you get more watchs, it grows **bigger**; if you take out some watchs, it **shrinks**.

### What is a slice in Go?------------

A slice is like a flexible array. Unlike regular arrays that have fixed size, slice can grow and shrink. they are more pawerful and flexible,allowing you to store multipal values of the same type in a single variable.

### How to use Slices with Examples:------------

1. ## Creating a slice directly
you can create a slice and put your items(like watch names) in it right away(turant)
```go
var watchs []string = []string{"sonata", "titen", "royel"}
fmt.Println(watchs)  //outputs: [sonata, titen, royel]
```

Shorthand:
```go
watchs := []string{"sonata", "titen", "royel"}
fmt.Println(watchs)  //outputs: [sonata, titen, royel]
```
Here, watchs is a slice holding three items

2. ## Creating Slice From An Array
if you have an array but you want the flexbility of a slice.
```go
array := [5]int{10,20,30,40,50}
slice := array[1:4]
fmt.Println(slice)  //outputs: [20,30,40]
```

this creates a slice from the second to the fourth element of the array.

3. ## Using the make function
Go provides a way to create slices using the make function.
```go
numbers := make([]int, 3)
numbers[0] = 1
numbers[1] = 2
numbers[2] = 3
fmt.Println(numbers)  //outputs: [1,2,3]
```

4. ## Adding items to a Slice
you can add more items to your slice using **append** function.
```go
 watchs := []string{"sonata", "titen"}
 watchs = append(watchs, "royel")
 fmt.Println(watchs)  //outputs: [sonata, titen, royel]
 ```
 Now, your watchs slice has three items.

 ### Types of Slices in Go.------------

 Slices in Go are versatile(bhut flexble) and can be categorized based on their creation method:
 1. ## Nil Slices
 A slice that hasn't been initialized yet:
```go
var numbers []int
if numbers == nil{
    fmt.Println(numbers is nil)
}
fmt.Println(numbers)  // Outputs: []
```
This slice has no elements and is **nil**

2. ## Empty Slices
An initialized slice with no elements:
```go
numbers := []int{}
fmt.Println(numbers) // Outputs: []
```

3. ## Slices with Values

Slices that are initialized with elements:
```go
numbers := []int{1,2,3}
fmt.Println(numbers)  // Outputs: [1,2,3]
```

4. ## Slices Created with make
Slices created using the **make** function, specifying **length** and **capacity**:
```go
//numbers := make([]type, length, capacity)
numbers := make([]int, 3,5)
fmt.Println(numbers) // Outputs: [0,0,0]
```
this creates a slice with a length of 3 and a capacity of 5
## The len() function returns the **length** of a slice (how many elements are present).
## The cap() function returns the **capacity** of a slice (how many elements can be stored).

##The copy() function is used to copy data from one slice to another.
```go
func copy(dst, src []Type) int
```
This function **returns the number of copied elements**, which is min(len(dst), len(src)).
```go
package main  
import "fmt"  

func main() {  
    values1 := []int{13, 52, 75, 96}  
    values2 := []int{43, 95, 86, 32, 76}  
    num1 := copy(values2, values1)  
    fmt.Println(values2)  
    fmt.Println("Number of elements copied:", num1)  
    num2 := copy(values2, values2[3:])  
    fmt.Println(values2)  
    fmt.Println("Number of elements copied:", num2)  
}
Output:
[13 52 75 96 76]  
Number of elements copied: 4  
[96 76 75 96 76]  
Number of elements copied: 2
```