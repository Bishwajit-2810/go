package functions

import "fmt"

func functions1(a, b int) int {
	return a + b

}
func functions2(a, b int) (result int) {
	result = a + b
	return

}
func functions3(a, b *int) {
	temp := *a
	*a = *b
	*b = temp

}

func Functions() {
	var a, b, c, d int
	fmt.Print("enter a number: ")
	fmt.Scan(&a)
	fmt.Print("enter a number: ")
	fmt.Scan(&b)
	fmt.Print("enter a number: ")
	fmt.Scan(&c)
	fmt.Print("enter a number: ")
	fmt.Scan(&d)

	fmt.Println("sum of first 2 num is: ", functions1(a, b))
	fmt.Println("sum of first 2 num is: ", functions2(a, b))
	fmt.Println("value of c and d is: ", c, d)
	functions3(&c, &d)
	fmt.Println("swapped value of c and d is: ", c, d)

}
