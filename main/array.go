package main 

import "fmt"

func main() {
   

var cars = []string{"Toyota","Ferrari","Ford","Marzd","Hynday"}
names := []string{"Pedro","José","Maria","Marcos"} 
var vL = [...]int{2,4,5,5,23,2422,232467,989}
cars[0] = "volvo"
life := []int{2,3,5,6}
// matrix 

life = append(life,12,34)
fmt.Println(life)
fmt.Println("Cap:",cap(life))
arr := append(cars, names...)
fmt.Println(arr)

mx1 := [5]int{}
mx2 := [5]int{1,2,3}
mx3 := [5]int{1,2,3,4,5}

fmt.Println(mx1)
fmt.Println(mx2)
fmt.Println(mx3)
leng := len(vL)
fmt.Println(leng)
fmt.Println("-------------------------------")
fmt.Println(cars)
fmt.Println(names)
fmt.Println(vL)

   
   
   
 
}