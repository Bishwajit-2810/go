package functions

import "fmt"

// _ is a write only variable which is ignored by the compailer

func function1(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

func ErrorHandling() {
	var a, b float64
	fmt.Print("enter a number: ")
	fmt.Scan(&a)
	fmt.Print("enter a number: ")
	fmt.Scan(&b)
	ans, err := function1(a, b)
	if err != nil {
		fmt.Println("error handling")
	} else {
		fmt.Println(ans)
	}

}
