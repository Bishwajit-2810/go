package syncwaitgroup

import (
	"fmt"
	"sync"
)

func printer(i int, wg *sync.WaitGroup) {
	fmt.Println("start", i)
	fmt.Println("end  ", i)
	defer wg.Done()
}

func SyncWaitGroup() {

	var wg sync.WaitGroup

	fmt.Println("programme started")

	for i := 1; i <= 5; i++ {
		wg.Add(1) //incremwnting web group count
		go printer(i, &wg)
	}
	wg.Wait()
	fmt.Println("programme ended")

}
