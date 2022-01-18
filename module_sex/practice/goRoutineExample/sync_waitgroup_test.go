package goRoutineExample

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Runner struct {
	Name string
}

func (r Runner) Run(wg *sync.WaitGroup) {
	fmt.Println(r.Name, "开始跑")
	rand.Seed(time.Now().UnixNano())

	time.Sleep(time.Duration(rand.Uint64()%10) * time.Second)
	fmt.Println(r.Name, "跑到终点")
}

func TestPlaySports(T *testing.T) {
	runnerCount := 10

	runners := []Runner{}

	wg := sync.WaitGroup{}
	wg.Add(runnerCount)

	for i:= 0; i<runnerCount;i++ {
		runners = append(runners, Runner{
			Name: fmt.Sprintf("%d", i),
		})
	}

	for _, runnerIter := range runners {
		fmt.Println(runnerIter)
	}
}
