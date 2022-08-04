package main

import "fmt"

func main() {
	type Student struct {
		name   string
		age    int
		course string
	}

	/*
		var name string
		var age int
		var course string
	*/
	var (
		name   string
		age    int
		course string
	)

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter your age:")
	fmt.Scanf("%i", &age)
	fmt.Print("Enter your course:")
	fmt.Scanf("%s", &course)

	std := Student{name, age, course}

	if age < 18 {

		fmt.Println(" Under age..")

	} else {

		fmt.Println("Name: ", std.name)
		fmt.Println("Age: ", std.age)
		fmt.Println("Course: ", std.course)

	}
}
