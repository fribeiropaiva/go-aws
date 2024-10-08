package main

import "fmt"

// Defining a struct
type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

// This is a receiver function
// It is a function that is attached to a struct
// It is like a method in other languages
// the asterisk is used to pass the struct by reference
// so the function can change the struct
func (p *Person) changeName(newName string) {
	p.Name = newName
}

func main() {
	myPerson := NewPerson("Filipe", 34)

	// This will print the struct without the property names
	fmt.Println(myPerson) // {Filipe 34}

	// This will print the struct with the property names
	fmt.Printf("This is my person %+v\n", myPerson) // This is my person {Name:Filipe Age:34}

	// This is how you use a receiver function
	// Now it can only be user by a struct of type Person
	myPerson.changeName("John")

	fmt.Println(myPerson)
	fmt.Println("------------")                             // {John 34}
	fmt.Printf("This is my changed person %+v\n", myPerson) // This is my person {Name:John Age:34}
}
