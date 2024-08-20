package mapsGo

import "fmt"

func MapsInGo() {
	// name <--> grade
	bkGrade := make(map[string]float64)
	bkGrade["bk 1st semester"] = 3.76
	bkGrade["bk 2nd semester"] = 3.51
	bkGrade["bk 3rd semester"] = 3.82
	bkGrade["bk 4th semester"] = 3.57
	bkGrade["bk 5th semester"] = 0.00

	fmt.Println("semester--------grade")
	for index, value := range bkGrade {
		fmt.Println(index, value)
	}
	delete(bkGrade, "bk 5th semester")
	fmt.Println("semester--------grade")
	for index, value := range bkGrade {
		fmt.Println(index, value)
	}
	semester, Exists := bkGrade["bk 4th semester"]
	fmt.Println(semester, Exists)
}

// unordered data
