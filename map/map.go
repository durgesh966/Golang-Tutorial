package main

import "fmt"

func main() {
	// ageOfStudent := make(map[string]int)

	// ageOfStudent["Durgesh"] = 21
	// ageOfStudent["piyush"] = 22
	// ageOfStudent["Aayush"] = 24
	// ageOfStudent["Bipul Kumar Dubey"] = 25

	// fmt.Println(ageOfStudent["Aayush"]) // 24

	// ageOfStudent["Aayush"] = 25
	// fmt.Println(ageOfStudent["Aayush"]) // 25

	// delete(ageOfStudent, "Aayush")
	// fmt.Println(ageOfStudent["Aayush"]) //0

	// Gread, Exist := ageOfStudent["Aayush"]
	// fmt.Println("the gread of Aayush", Gread) // return Gread num else 0
	// fmt.Println("Exist", Exist)  // true if present else false

	// for index, element := range ageOfStudent {
	// 	fmt.Printf("The name of Student %s and his age is %d\n", index, element)
	// }

	// ----------------- Output ------------------------
	// 	The name of Student piyush and his age is 22
	// The name of Student Aayush and his age is 25
	// The name of Student Bipul Kumar Dubey and his age is 25
	// The name of Student Durgesh and his age is 21

	person := map[string]int{
		"ayush":      25,
		"bipul":      24,
		"raghvendra": 23,
		"mayank":     22,
		"durgesh":    21,
	}

	for name, age := range person {
		fmt.Printf("The Name of Student %s and there age is %d\n", name, age)
	}

	// ---------------- Output -------------------
	// The Name of Student raghvendra and there age is 23
	// The Name of Student mayank and there age is 22
	// The Name of Student durgesh and there age is 21
	// The Name of Student ayush and there age is 25
	// The Name of Student bipul and there age is 24
}
