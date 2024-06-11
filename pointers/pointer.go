package main
import "fmt"

func main(){
	var age int = 20
	var p *int
	p = &age

	fmt.Println("age is", *p);
}