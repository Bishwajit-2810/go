package structure

import "fmt"

type Person struct {
	FirstName string
	Lastname  string
	Age       int
}

func Structure() {
	var person1 Person
	person1.FirstName = "Bishwajit"
	person1.Lastname = "Chakraborty"
	person1.Age = 22
	fmt.Println(person1)

	var person2 Person = Person{
		FirstName: "bk",
		Lastname:  "c",
		Age:       22,
	}
	fmt.Println(person2)

}

// same thing as C
