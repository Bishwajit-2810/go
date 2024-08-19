package array

import "fmt"

func Array() {
	var arr [5]string
	arr[0] = "bcd"
	arr[1] = "cd"
	arr[2] = "d"
	arr[3] = "abcd"

	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

}

// default values of array
//  int ----> 0
//  float---> 0.00
//  string --> ""
//  bool-----> false
//  pointer--> nil
