package variables

import "fmt"

var a int = 100
var b float64 = 100.55
var c string = "hello"
var d bool = false
var e = "world"

const pi = 3.1416

// name := "Bishwajit"
func Variable() {
	f := "Bishwajit" //cann't assign outside of function
	fmt.Println(a, b, c, d, e, f, pi)
	fmt.Printf("%d %.2f %s %t %s %s %.3f \n", a, b, c, d, e, f, pi)
}

// name capital start means punlic
// name non capital start means private
// %q print with quatation
