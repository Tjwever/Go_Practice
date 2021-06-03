package main

import (
	"fmt"
)

func main() {
	//Below is just a few example of me playing around with Go syntax
	var ages = [5]int{31, 22, 65, 13, 49}

	// the := is like declaring a variable inside a function.  This doesn't work outside of a function
	name := "Tim"
	age := 36
	fmt.Println("Sup, Dudes?, my name is " + name)
	fmt.Println("I am", age)
	fmt.Println("And this is a bunch of numbers:")
	fmt.Println(ages, len(ages))

	fmt.Println(" ")
	fmt.Println("-----------------------")
	fmt.Println("Creating Books Below")
	fmt.Println(" ")

	//Below, we're going to create some book "objects" that relate to the info provided above
	book1 := Book{"The Dark Lagoon", "Some Person", 343}
	//And since Go doesn't like it if you declare/create something without using it, we'll use it below
	book1.print_details()
	//Just for another example, I'll go ahead and create one more Book
	fmt.Println(" ")
	book2 := Book{"The Light Lagoon", "Another Person", 434}
	book2.print_details()

	test()

}
