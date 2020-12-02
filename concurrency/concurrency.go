package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func UploadFile(wg *sync.WaitGroup, i int) {
	fmt.Println(i, "is start")
	time.Sleep(time.Millisecond * 3)
	fmt.Println(i, "is done")
	wg.Done()
}

const NubGoroutines = 26

func Service() {
	var wg sync.WaitGroup
	wg.Add(NubGoroutines)
	for i := 0; i < NubGoroutines; i++ {
		go UploadFile(&wg, i)
	}
	wg.Wait()
}
