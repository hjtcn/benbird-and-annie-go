package common

import (
	"testing"
	"time"
)

func TestFunc(t *testing.T) {
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		t.Log("运行时间：", finishTime.Sub(startTime))
	}()
	arr := make([]int, 1, 100)
	t.Log(arr, len(arr), cap(arr))

	arr1 := []int{}
	res := copy(arr1, arr)

	t.Log(res)

	arr = append(arr, 2)
	t.Log(arr, len(arr), cap(arr))

	t.Log(arr1, len(arr1), cap(arr1))
}
