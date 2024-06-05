package main

import "fmt"

func main() {
	day := 10

	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednusday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("suterday")
	case 7:
		fmt.Println("sunday")
	default:
		fmt.Println("in a week have only 7 days")
	}
}
