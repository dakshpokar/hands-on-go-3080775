// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func getVal() string {
	return val
}

func updateVal() string {
	val = "OMG"
	return val
}

func main() {
	// create a local variable "val" and assign it an int value
	var val int = 4
	// print the value of the local variable "val"
	fmt.Println(val)

	// print the value of the package-level variable "val"
	fmt.Println(getVal())

	// update the package-level variable "val" and print it again
	fmt.Println(updateVal())

	// print the pointer address of the local variable "val"
	fmt.Println(&val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 36
	fmt.Println(val)
}
