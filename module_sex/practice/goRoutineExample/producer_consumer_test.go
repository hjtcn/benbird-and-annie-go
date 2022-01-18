package goRoutineExample

import (
	"fmt"
	"sync"
	"testing"
)

type Store struct {
	DataCount int
	Max       int
	lock      sync.Mutex
}

type Producer struct {
}

func (Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.DataCount == s.Max {
		fmt.Println("库存已满，不生产")
		return
	}

	fmt.Println("开始生产 +1")
	s.DataCount++
}

type Consumer struct {
}

func (Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.DataCount == 0 {
		fmt.Println("库存为空，不消费")
		return
	}

	fmt.Println("开始消费 -1")
	s.DataCount--
}

func TestWorker(t *testing.T) {
	s := &Store{
		Max: 10,
	}

	pCount, cCount := 50, 50

	for i := 0; i < pCount; i++ {
		go func() {
			Producer{}.Produce(s)
		}()
	}

	for j := 0; j < cCount; j++ {
		go func() {
			Consumer{}.Consume(s)
		}()
	}
}
