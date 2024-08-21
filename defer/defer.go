package defer_go

import "fmt"

func Defer_go() {
	fmt.Println("first line")
	fmt.Println("secound line")
	defer fmt.Println("last line 1")
	defer fmt.Println("last line 2")
	fmt.Println("third line")

}

// defer runs in the last of the programme
// use stack if there are many defer
// follows lifo
// mostly for function call
