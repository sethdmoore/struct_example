package main

import (
	"fmt"
	"github.com/sethdmoore/struct_example/bug"
)

func main() {
	myBug := bug.NewBug("Jimmy")
	// I can access the number of legs
	fmt.Printf("%s: Number of legs: %d\n", myBug.Name, myBug.Legs)

	//Bug loses a leg
	myBug.Legs -= 1

	fmt.Printf("%s now has %d legs\n", myBug.Name, myBug.Legs)

	// If we try to set the line below...
	// myBug.killable = true

	// A) There are no bugs with godmode allowed
	// B) This will cause the following error:
	// ./main.go:20:7: myBug.killable undefined (cannot refer to unexported field or method killable)

	// We exposed a method called 'IsKillable'
	// Remember, Capitalized words in Golang are Public / exported!
	if myBug.IsKillable() {
		// This works, and is a safe way to refer to the 'killable' boolean of
		// myBug, without allowing it to be modified from this package
		fmt.Printf("%s is killable\n", myBug.Name)
	}

	// All of these safety features prevent us from creating a GOD MODE BUG
	// The end
}
