package Intermediate

import "fmt"

func main() {
	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
	fmt.Println(sequence2())
	fmt.Println(sequence2())
	fmt.Println(sequence2())
	fmt.Println(sequence2())

	subtracter := func ()func(int)int  {
		countdown := 99
		return func(i int) int {
			countdown -= i
			return countdown
		}
	}()
	fmt.Println(subtracter(99))
}

func adder() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

