package timedate

import (
	"fmt"
	"time"
)

func TimeDate() {
	var curtime = time.Now()
	fmt.Println(curtime)
	fmt.Printf("%T \n ", curtime)
	formatted := curtime.Format("02-01-2006, 15:04:05")
	println(formatted)
	formatted = curtime.Format("02-01-2006, 3:04:05 PM")
	println(formatted)
	formatted = curtime.Format("02/01/2006, 3:04:05 PM, Monday")
	println(formatted)

	layoutStr := "02-01-2006"
	dateStr := "21-08-2024"
	formattedTime, _ := time.Parse(layoutStr, dateStr)
	fmt.Println(formattedTime)
	fmt.Printf("%T \n", formattedTime)

	// add one more day in current time
	newTime := curtime.Add(24 * time.Hour)
	fmt.Println(newTime.Format("02/01/2006, 3:04:05 PM, Monday"))
}

// 2006-01-02 15:04:05
// January 2,2006, 3:04:05 PM Monday
// dd-mm-yyyy doesnot supported by go
