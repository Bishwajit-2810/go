package goroutine

import (
	"fmt"
	"time"
)

func goodMorning() {
	fmt.Println("good morning")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("completed")
}
func goodNight() {
	fmt.Println("good night")
}

func GoRoutine() {
	fmt.Println("started")
	go goodMorning()
	go goodNight()
	time.Sleep(2000 * time.Millisecond)
}
