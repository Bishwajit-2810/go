package pointers

import "fmt"

func Pointer() {
	var number int = 2
	var ptr *int = &number
	fmt.Println(number, &number, *ptr, ptr)
	num := 5
	ptr2 := &num
	fmt.Println(num, &num, *ptr2, ptr2)

}

// default value of pointer is nil
// havily used for modify by reference (functions)
