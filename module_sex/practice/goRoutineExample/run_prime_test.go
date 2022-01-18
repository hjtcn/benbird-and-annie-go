package goRoutineExample

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func collectPrime(startNum, endNum int) []int {
	var res []int

	for i := startNum; i <= endNum; i++ {
		if i == 1 || i == 2 {
			continue
		}

		for j := 2; j < i; j++ {
			if i % j != 0 {
				break
			}
		}

		res = append(res, i)
	}

	return res
}

func TestRunPrime(t *testing.T) {
	startTiem := time.Now()
	result := []int{}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		result = append(result, collectPrime(2, 100000)...)
	}()

	go func() {
		defer wg.Done()
		result = append(result, collectPrime(100001, 200000)...)
	}()

	wg.Wait()
	finishTime := time.Now()

	fmt.Println(len(result))
	fmt.Println("共耗时:", finishTime.Sub(startTiem))
}
