package conditionals

import "fmt"

func Conditional() {
	tax := 25
	if tax > 30 {
		fmt.Println("tax is higher than 30")

	} else {
		fmt.Println("tax is lesser than 30")
	}

	switch i := 5; { // i gets value of 5
	case i <= 2:
		fmt.Println("i is less than or equal to two")
	case i <= 4:
		fmt.Println("i is between three and five")
	default:
		fmt.Println("i is greater than five")
	}

}
