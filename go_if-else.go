//  for test the conditions

package main 
import "fmt"

func main(){
	var a int = 11
	if(a%2==0){
		fmt.Println("this is even number")
	}else if(a==0){
		fmt.Println("this is zero")
	}else{
		fmt.Println("this is odd number")
	}
}