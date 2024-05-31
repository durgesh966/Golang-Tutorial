package main 
import "fmt"

func divide(a, b float64) (float64, error){
	if(b == 0){
		return 0, fmt.Errorf("the number shuld be greater than zero");
	}else{
		return a/b, nil
	}
}

func main(){
	fmt.Println("division started");

	// ------------- handle error -----------------------
	// ans, err := divide(10, 5);
	// if(err != nil){
	// 	fmt.Println("error handling");
	// }
	// fmt.Println("output of division is", ans);


	ans, _ := divide(10, 5);
	fmt.Println("output of division is", ans);
}