package main

import "fmt"

func main() {

  
 for i:=0;i <=10;i++ {
	fmt.Println("Count:",i)
 }
	fmt.Print("Enter something: ")
	var input float64
	fmt.Scanf("%f", &input)

	if input < 0 {
       fmt.Println("Oh !!!!!!!!!!!")
	}else {
		
	  output := input * 2
	  fmt.Println(output)
	}


	

}
