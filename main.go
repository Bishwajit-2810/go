package main

import (
	"fmt"
	"learning_go/array"
	"learning_go/conditionals"
	"learning_go/functions"
	"learning_go/loop"
	mapsGo "learning_go/maps"
	"learning_go/myutill"
	"learning_go/pointers"
	"learning_go/structure"
	"learning_go/userinput"
	"learning_go/variables"
)

func main() {

	fmt.Println("hello world")
	myutill.PrintMessage()
	variables.Variable()
	userinput.UserInput()
	functions.Functions()
	functions.ErrorHandling()
	array.Array()
	array.Slice()
	conditionals.Conditional()
	loop.Loop()
	mapsGo.MapsInGo()
	structure.Structure()
	pointers.Pointer()

}
