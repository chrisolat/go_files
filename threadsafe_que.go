package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	size     int
	capacity int
	buf      []interface{}
	mtx      sync.Mutex
	cond     *sync.Cond
}

func newQue(cap int) *Queue {
	que := &Queue{
		capacity: cap,
		buf:      []interface{}{},
	}
	que.cond = sync.NewCond(&que.mtx)
	return que
}

func (que *Queue) push(val interface{}) error {
	que.mtx.Lock()
	defer que.mtx.Unlock()
	// Wait until there is space in the queue
	for que.size >= que.capacity {
		que.cond.Wait()
	}
	// Add the item to the queue
	que.buf = append(que.buf, val)
	que.size++
	fmt.Println("Pushed:", val, "Queue size:", que.size)
	// Signal one waiting consumer that an item is available
	que.cond.Signal()
	return nil
}

func (que *Queue) pop() (interface{}, error) {
	que.mtx.Lock()
	defer que.mtx.Unlock()
	// Wait until there is an item in the queue
	for que.size <= 0 {
		que.cond.Wait()
	}
	// Remove the item from the queue
	item := que.buf[0]
	que.buf = que.buf[1:]
	que.size--
	fmt.Println("Popped:", item, "Queue size:", que.size)
	// Signal one waiting producer that space is available
	que.cond.Signal()
	return item, nil
}

func main() {
	var wg sync.WaitGroup
	que := newQue(5) // Initialize queue with capacity 5
	wg.Add(1)
	// Producer goroutine
	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			que.push(i)
		}
		wg.Done()
	}()

	// Consumer goroutine
	go func() {
		for i := 0; i < 10; i++ {
			item, err := que.pop()
			if err != nil {
				fmt.Println("Error popping item:", err)
			} else {
				fmt.Println("Popped item:", item)
			}
			wg.Done()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final queue size:", que.size)
	// select {}
}
