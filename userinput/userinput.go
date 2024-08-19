package userinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UserInput() {

	var a int

	fmt.Print("enter name: ")
	var name string
	fmt.Scan(&name)
	fmt.Print("enter a number: ")
	fmt.Scan(&a)
	fmt.Print("enter full address: ")
	reader := bufio.NewReader(os.Stdin)
	testString, _ := reader.ReadString('\n')   // it also takes the next line
	testString = strings.TrimSpace(testString) // it's for removing the extra line
	fmt.Println(a, name, testString)

}
