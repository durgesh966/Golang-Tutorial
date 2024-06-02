package main

import "fmt"

func main() {
	fmt.Println("hello this is a slice")
	// number := []int{1, 2, 3, 4, 5}
	// number = append(number, 6, 7, 8, 9)
	// fmt.Println("the number is", number)

	// ---------------- crate using make funk ------------------

	numbers := make([]int, 0)
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("number is", numbers)
	fmt.Println("length is", len(numbers))
	fmt.Println("capacity is", cap(numbers))

}
