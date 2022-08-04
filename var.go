package main

import (
	"fmt"
	"math"
) 

func main() {

   type person struct {
	 name string
	 age int 
	 height float32 
   }



   p := person{"pedro",21,1.80}
    
	var name, age = "Pedro" , 21
	var married bool = false 
	var numberBig float32 = 174974937594594549859
	var numberInf  float64 = 4940737994774084804808384970370
	const bornY = 2001
	const currentY int = 2022
	var dog string = "Max"
	ageY := currentY - bornY 
	x := 0
	fmt.Println("Sqrt(49) =",math.Sqrt(49))
	fmt.Println("Person name:",p.name)
	fmt.Println("Person age:",p.age)
	fmt.Println("Person height:",p.height)
	increment := func() int {
		    x++
			return x
	}

	var (
		a  int = 2
		b  int = 3
		c  string = "Hello"
	)
    
	fmt.Println("---- var ------")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	
    fmt.Println("increment: ",increment())

    // fmt.Println()
	/**
	  int 
	  float32, float64
	  bool 
	  string 
	*/

	fmt.Println("My name is ",name, "and I´m ",age,"years old")
	fmt.Println("You are married:",married)
	fmt.Println("Big number is:",numberBig)
	fmt.Println("Big number is:",numberInf)
	fmt.Println("My age is:",ageY)
	fmt.Println("My dog is ",dog)

}