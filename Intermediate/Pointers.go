package main

import "fmt"

func main() {
	var ptr *int  // we use *(datatype) to mention that is a pointer. The pointers have their own data type (int, string, float, etc)
	var a int = 10
	ptr = &a //referencing

	fmt.Println(a)
	fmt.Print(ptr)
	fmt.Println(&ptr)//dereferencing a pointer --we can't do aritmetic operations using pointers(Golang). we have 2 actions -- referencing an dereferencing
	//the zero value is nil --> pointers

	modifyValue(ptr)
	fmt.Println(a)
	fmt.Println(*ptr)
}

// we are using *int to indicate that is a pointer variable and NOT int. If we use * for second time in a pointer, we get the real value 
func modifyValue(ptr *int){
	*ptr++
}