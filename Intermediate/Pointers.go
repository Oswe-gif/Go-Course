package main

import "fmt"

func main() {
	var ptr *int
	var a int = 10
	ptr = &a //referencing

	fmt.Println(a)
	fmt.Print(ptr)
	fmt.Println(&ptr)//dereferencing a pointer --we can't do aritmetic operations using pointers(Golang). we have 2 actions -- referencing an dereferencing
}