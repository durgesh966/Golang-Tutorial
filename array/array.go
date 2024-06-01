package main

import "fmt"

func main(){
	// var array = [5] int {1,3,4,5,6}
	// fmt.Println("length of array is", len(array));

	var name [4] string
    name[0] = "Durgesh"
	name[1] = "yuvraj"
	name[2] = "harsh"
	name[3] = "Bipul"
	fmt.Println("the array of name is", name);
	// fmt.Println("the array of name is", len(name));
	fmt.Println("1 name", name[1]);
}