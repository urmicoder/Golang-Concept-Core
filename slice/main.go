package main

import "fmt"

func main() {
	array := [5]int{10, 20, 30, 40, 50}
	//slice := array[1:4]
	slice := array[:]
	slice = append(slice, 60)
	fmt.Println(slice) // Outputs: [10 20 30 40 50 60]

}
