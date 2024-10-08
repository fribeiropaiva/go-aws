package main

import (
	"fmt"
	"slices" // helper for working with slices
)

func main() {
	// Arrays in Go are fixed size
	// you can not have mixed types in an array
	animals := [2]string{}

	// Or you can do it like this
	animals2 := [2]string{"Dog", "Cat"}

	// This is how you assign values to an array
	animals[0] = "Dog"
	animals[1] = "Cat"

	fmt.Println(animals)
	fmt.Println(animals2)

	// Slices are like arrays but they are dynamic
	// you can add and remove elements from a slice
	// you can also have mixed types in a slice
	sliceAnimals := []string{"Dog", "Cat"}

	// This is how you add an element to a slice
	sliceAnimals = append(sliceAnimals, "Bird")

	sliceAnimals = slices.Delete(sliceAnimals, 0, 1)

	fmt.Println(sliceAnimals)

	// Regular for loop
	for i := 0; i < len(sliceAnimals); i++ {
		fmt.Printf("This is my animal %s\n", sliceAnimals[i])
	}

	// Using range function
	for index, value := range sliceAnimals {
		fmt.Printf("This is my animal %s at index %d\n", value, index)
	}

	for value := range 10 {
		fmt.Println(value)
	}

	// There is no while loop in Go
	// But you can use a for loop like this
	// to create a while loop
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}
}
