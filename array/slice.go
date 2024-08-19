package array

import "fmt"

func Slice() {
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers, len(numbers))
	numbers = append(numbers, 6, 7, 8, 9)
	fmt.Println(numbers, len(numbers))

	number := make([]int, 3, 5)
	number = append(number, 6, 7, 9)
	fmt.Println("Slice:", number)
	fmt.Println("Length:", len(number))
	fmt.Println("Capacity:", cap(number))
	number = append(number, 6, 7, 8, 9)
	fmt.Println("Slice:", number)
	fmt.Println("Length:", len(number))
	fmt.Println("Capacity:", cap(number))
	number = append(number, 6, 7, 8, 9, 10, 11)
	fmt.Println("Slice:", number)
	fmt.Println("Length:", len(number))
	fmt.Println("Capacity:", cap(number))
}

//in slice
// lenght is the total element number
// capacity is how much is possible to store(increase in 2*previas capacity)
// when length > capacity capacity increases
