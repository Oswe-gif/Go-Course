package main

import (
	"errors"
	"fmt"
)

var generalValue = "jeje"
func main(){
	fmt.Println(generalValue)
	generalValue = "what?"
	fmt.Println(generalValue)
	generalValue := "blash"
	fmt.Println(generalValue)

	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}

	number := []int{2,3,5,6}

	for index, value := range number {
		fmt.Printf(" Index: %d, value: %d",index,value)
	}

	println()
	for true {
		fmt.Print("ehhhhhh")
		break
	}
	println()
	//source := rand.NewSource(time.Now().UnixNano())
	//random := rand.New(source);
	//target := random.Intn(100)+1
	//var guess int

	//for guess != target {
	//	fmt.Println("enter your guess")
	//	fmt.Scanln(&guess)
//
//		if guess == target {
	//		fmt.Println("Well done")
	//		break;
	//	}
	//	if guess > target{
	//		fmt.Println("The number is smaller")
	//	}
	//	if guess < target{
	//		fmt.Println("The number is higher")
	//	}
	//}

	matrix := [3][3]int{{1,2,3},{2,3,4},{1,2,3}}
	fmt.Println(matrix)

	var matrix2 [3][3]int= [3][3]int{{1,2,3},{2,3,4},{1,2,3}}
	fmt.Println(matrix2)

	// just copy the value
	copiedValue := matrix;
	copiedValue[0][0] = 1
	fmt.Println(copiedValue)

	//change the originalMatrix

	var addressMatrix *[3][3]int = &matrix2
	(*addressMatrix)[0][0] = -1
	fmt.Println(matrix2)


	//map

	var myMap map[string]int
	fmt.Println(myMap)

	myMap = map[string]int{}
	myMap["key1"]=1
	fmt.Println(myMap)

	m := make(map[string]int) 

	fmt.Println(m)

	//range

	for i,v := range "Hello world!"{
		fmt.Println("index: ",i, " Rune: ",v)
		fmt.Printf("index: %d, Rune %c \n",i,v)
	}

	//for _,v := range "Hello world!"{
	//	fmt.Println(" value: ",v)
	//}

	//slices

	slice := make([]int, 3)
	fmt.Println(slice)
	fmt.Println(slice == nil)
	var s []int           // declarado pero no inicializado
	fmt.Println(s == nil)
	slice = append(slice, 1)
	fmt.Println(slice)

	//it was modified the real array using slice
	fmt.Println(number)
	secondSlice := number[0:3]
	secondSlice[0]=-100
	fmt.Println(number)

	//but if I change the size, the pointer changes too, because does not match array's elements
	// if we want to keep the pointer, we have to ensure that array's size and slice's size match

	secondSlice = append(secondSlice, 3)
	secondSlice[0]=100
	fmt.Println(number)

	secondSlice = append(secondSlice, 3)
	secondSlice[0]=-2200
	fmt.Println(number)

	//anonymous function
	func ()  {
		fmt.Println("calling")
	}()

	greet := func ()  {
		fmt.Println("calling")
	}
	greet()
	//general functions
	operation := add
	fmt.Println(operation(1,2))

	fmt.Println(applyOperation(2,2,add))

	functionInsideOfFunction := createMultiplier(3)
	fmt.Println(functionInsideOfFunction(2))
	fmt.Println(createMultiplier(2)(2))

	result, error := compare(1,2)
	if error == nil {
		fmt.Println(result)
	}

	fmt.Println(generalFunction(3,4))
	fmt.Println(generalFunctionX2(3,4))
	fmt.Println(sum("sum",1,2,3))
}

func add(a,b int)int{
	return a+b
}

// a function that takes a function as an argument
func applyOperation(a int, b int, operation func(int, int)(int))int{
	return operation(a,b)
}

//function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}

func compare(a,b int)(string, error){
	if a > b {
		return "a is greater than b", nil
	}else if b > a{
		return "b is greater than a", nil
	}else{
		return "", errors.New("man, they are equals")
	}
}

//a new way
func generalFunction(a,b int)(x,y int){
	result1 := a+b
	result2 := a-b
	return result1, result2
}
func generalFunctionX2(a,b int)(x,y int){
	x = a+b
	y = a-b
	return
}

func sum(returnString string, nums ...int)(string, int){
	total := 0
	for _, v := range nums{
		total += v
	}
	return returnString,total
}


