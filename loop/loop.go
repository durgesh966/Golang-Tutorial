package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("num", i)
	// }

	// ------------------ infinity loop -------------------------
	// count := 0
	// for {
	// 	fmt.Println("hello durgesh bhai", count)
	// 	count++
	// 	if count == 5 {
	// 		break
	// 	}
	// }

	// -------------- use range in array ----------------------
	// arr := []int{2, 3, 4, 5, 56, 6}
	// for pos, num:= range arr {
	// 	fmt.Println("num and pos are", num, pos)
	// }

	// -------------- use range in string ---------------------
	data := "hello, Durgesh"
	for index, value := range data {
		fmt.Printf("the position and, %d character in strings are %c\n", index, value)
	}
}
