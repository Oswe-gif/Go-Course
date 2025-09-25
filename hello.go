package main

import (
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

	// just a copy value
	copiedValue := matrix;
	copiedValue[0][0] = 1
	fmt.Println(copiedValue)

	//change the originalMatrix

	var addressMatrix *[3][3]int = &matrix2
	(*addressMatrix)[0][0] = -1
	fmt.Println(matrix2)

}

