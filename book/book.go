package book

import "fmt"

// To illustrate the concept of Go's "OOP", we'll use a book in this example
// We'll start by defining the struct

type Book struct {
	//Inside we'll create the properties of what a book would contain
	name   string
	author string
	pages  int
}

//Here we're going to create a 'method', that will contain some info on what we want to do with the Book 'Object'
func (b Book) print_details() {
	fmt.Printf("Book %s was written by %s.", b.name, b.author)
	fmt.Printf("\nIt contains %d pages. \n", b.pages)
}