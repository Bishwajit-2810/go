package stringsgo

import (
	"fmt"
	"strings"
)

func Strings() {
	str := "welcome,this a go learning repository"
	parts := strings.Split(str, ",")
	fmt.Println(parts[0])
	fmt.Println(parts[1])
	fmt.Println(parts)

	count := strings.Count(str, "e")
	fmt.Println(count)

	str = "       welcome      ,    this a go learning repository       "
	newStr := strings.TrimSpace(str)
	fmt.Println(newStr)

	str = "welcome,"
	newStr = "this a go learning repository"
	joinedStr := strings.Join([]string{str, newStr}, " ")
	fmt.Println(joinedStr)
}
