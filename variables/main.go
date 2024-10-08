package main

import "fmt"

// Go applications can have multiple entry points (multiple microservices)
// as long as they are separated by their own folder, they act as their own
// entry point.
func main() {
	// Explicit variable declaration
	var myName string = "Filipe"

	// This is implicit variable declaration
	// type is inferred here
	myAge := 34

	// This is the way to do concatenation in Go
	// %s is a placeholder for a string
	fmt.Printf("Hello, my name is %s, my age is %d\n", myName, myAge)

	// Go has the concept of zero values
	// These are the default values for each type
	// string: ""
	// int: 0
	// bool: false
	// float: 0.0
	// complex: 0+0i
	// pointer: nil
	// function: nil
	// interface: nil
	var myFriendsName string // ""
	var myBool bool          // false
	var myOtherInt int       // 0

	fmt.Printf("My friend's name is %s, he is %d years old, is he a good friend? %t\n", myFriendsName, myOtherInt, myBool)
}
