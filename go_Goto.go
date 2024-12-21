package main
import "fmt"

func main(){

	var input int
    var chance int = 5
	Dhanraj:
	fmt.Println("Guess the Number you have only ", chance)
	fmt.Scanln(&input)
    if(input == 34){
		fmt.Println("you got it ")
	}else if(chance ==1){
          fmt.Println("your lose")
	}else{
		chance--
		goto Dhanraj
	}
}