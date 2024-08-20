package loop

import "fmt"

func Loop() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("5 skipped")
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()
	i := 0
	for {
		if i == 15 {
			fmt.Println(" infinite loop breaked")
			break
		}
		i++
	}
	fmt.Println()
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println("index value")
	for index, value := range numbers {
		fmt.Println(index, "    ", value)
	}
}
