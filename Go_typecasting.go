package main
import( 
      "fmt"
      "strconv" 
	)

func main(){
   var a int =29
   var b float64 = 223.33
   fmt.Println(float64(a))
   fmt.Println(int(b))

   var c string = "102"
   var d string = "102.78"
    newint, _ := strconv.ParseInt(c,0,64)
    newFloat, _ := strconv.ParseFloat(d,64)

   fmt.Println(newint)
   fmt.Println(newFloat)

}
