package main

import "fmt"

func main() {
	fmt.Println("this is a ef-else tutes")

	x := 6

	if x > 5 {
		fmt.Println("x is grater than 5")
	} else {
		fmt.Println("x is smaller then 5")
	}

	y := 25

	if x > 5 && (y < 10 || y > 10) {
		fmt.Println("yah task completed")
	} else {
		fmt.Println("try another way")
	}
}
