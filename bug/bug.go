package bug

// This package scope variable (similar to "global variable") is **lowercase**
// Therefore it cannot be accessed outside of the scope of this package
var defaultKillable = true

// This struct (similar to class or object) is **lowercase**
// therefore it is private and inaccessible outside of this package
// The reason this is private is because we don't want GOD MODE bugs in
// our example
type bug struct {
	// Golang treats Capitalized variables as **exported**
	// Exported is analagous to Public in C# (Unity)
	Name string
	Legs int

	// Golang does not export **lowercase** variables.
	// It is equivalent to a private variable
	killable bool
}

// This function is **lowercase**
// This makes it private or unexported
// You cannot call this function outside the package
func setDefaultLegs() int {
	return 6
}

// NewBug creates a new bug struct and returns it
// Structs may have methods attached, sometimes they do not
func NewBug(name string) bug {
	// instantiate the struct ("class without methods")
	thisBug := bug{}

	// This will be set to false, referring to the package scope
	thisBug.killable = defaultKillable

	// I'm only using setDefaultLegs to show that you can call internal
	// package-scoped function from here. You cannot call setDefaultLegs()
	// outside of this package.
	thisBug.Legs = setDefaultLegs()

	// This will set the name when we call this function
	thisBug.Name = name

	// This will return a new struct with defaults set
	return thisBug
}

// Capitalized == Exported == Public method!
// Attach a method to the struct
// This method can be called outside this package inside of main.go
// We can check the internal variables of the struct
// Attaching this method to the struct makes it be able to be called
// like myBug.IsKillable()
func (b *bug) IsKillable() bool {
	return b.killable
}
