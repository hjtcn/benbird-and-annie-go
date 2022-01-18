package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		fmt.Println(time.Now().Sub(startTime))
	}()
	totalNum := 0
	totalNumLock := sync.Mutex{}
	totalWorkers := 100

	wg := sync.WaitGroup{}
	wg.Add(totalWorkers)

	for i := 0; i < totalWorkers; i++ {
		go func() {
			defer wg.Done()
			totalNumLock.Lock()
			totalNum += 100
			totalNumLock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("总字数：", totalNum)
}
