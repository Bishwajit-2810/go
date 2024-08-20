package dataconversion

import (
	"fmt"
	"strconv"
)

func Dataconversion() {
	var num int = 566
	fmt.Print(num)
	fmt.Printf(" %T\n", num)

	var data float64 = float64(num)
	fmt.Print(data)
	fmt.Printf(" %T\n", data)

	str := strconv.Itoa(int(num))
	fmt.Print(str)
	fmt.Printf(" %T\n", str)

	number, _ := strconv.Atoi(str)
	fmt.Print(number)
	fmt.Printf(" %T\n", number)

	number2, _ := strconv.ParseFloat(str, 64)
	fmt.Print(number2)
	fmt.Printf(" %T\n", number2)

}
