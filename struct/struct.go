package main
import "fmt"

type Person struct{
	name string
	email string 
	age int
}

func main(){
   var pers1 Person
   var pers2 Person


   pers1.name = "Durgesh Bisen"
   pers1.email = "webdev.durgesh@gmail.com"
   pers1.age = 21

   pers2.name = "Bipul Kumar Dubey"
   pers2.email = "bipul@gmail.com"
   pers2.age = 24

//    fmt.Println("person 1 info is", pers1)
//    fmt.Println("person 2 info is", pers2)

   printFunction(pers1);
   printFunction(pers2)
}

func printFunction(pers Person){
    fmt.Println("person 1 info is", pers)
//    fmt.Println("person 2 info is", Person)

}

// person 1 info is {Durgesh Bisen webdev.durgesh@gmail.com 21}
// person 2 info is {Bipul Kumar Dubey bipul@gmail.com 24}