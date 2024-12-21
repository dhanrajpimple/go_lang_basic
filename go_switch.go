package main
import (
	"fmt"
)


func main(){
	fmt.Print("enter the number")

	var input int

	fmt.Scanln(&input)

	switch input {
	case 10:
		fmt.Printf("this is ten")
	case 20: 
	    fmt.Print("this is twenty")
	case 30:
		fmt.Printf("this is thirty")
	default:
		fmt.Printf("i don't know this input")
	}
}