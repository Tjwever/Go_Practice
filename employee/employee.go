package employee

import "fmt"

//Employee Section

type Employee struct {
	id int
	name string
	salary float32
}

//We can do a 'is-a' relationship here, that's kind of like ruby/rails wait of doing a 
//'has_many' or 'uses'
//Grouping this section together to show example
type Manager struct {
	Employee		//embedding a nameless employee
	managedSector string
}
func (e Employee) printNameMethod() {
	fmt.Println("Name: ", e.name)
}
func printNameFunction(e Employee) {
	fmt.Println("Name: ", e.name)
}
func test() {
	//Here, you are creating a Manager object, and passing in the info needed
	//for Employee which is id, name, and salary, then for Manager which is
	//the Employee struct data type and managedSector
	m := Manager{Employee{1, "Bob Dole", 1000.0}, "Engineering"}

	fmt.Println("---------Testing Manager Stuff-----------")
	//Then to test
	fmt.Println(m.id)		//this will print the managers id.  Or m.Employee.id

	//Here is just showing you the two different ways you could print this
	printNameFunction(m.Employee)
	m.printNameMethod()
	fmt.Println("---------Testing Manager Stuff End-----------")

}
//End Section

//Again, we are dipping into the Employee struct.  The reason for the * is that 
//this function in particular is going to mutate an instance of the type Employee
func (e *Employee) raiseSalary(bonus float32) {
	e.salary += bonus
}

//And since nothing is being changed or mutated here, we don't need the *
func (e Employee) formatter() string {
	return fmt.Sprintf("Id: %d, Name: %s, Salary: %.2f", e.id, e.name, e.salary)
}

//If you create a function and don't have a return, we have to make sure that this
//gets used somewhere.  Otherwise, it'll cry about not using it, unlike the formatter func
func (e Employee) Greeting() {
	fmt.Sprintf("Hello: %s", e.name)
}