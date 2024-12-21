package main
import "fmt"


func main(){

	num := []int{1,2,3,3,4,4,5,6,7}
	var sum int
	for _, v := range num{
		sum+=v
	}
	fmt.Println(sum)
	text := map[string]int{"dhanraj":1, "pimpole":2, "satara":3, "Nashik":4}

	for f, s := range text {
		fmt.Printf("%s %d\n", f,s)
	}
	for f :=range text {
		fmt.Println(f)
	}
	for _,f :=range text {
		fmt.Println(f)
	}
	for i, c := range "best software developer"{
		fmt.Println(i, c)
		fmt.Printf("%d %c\n", i, c)
	}
}