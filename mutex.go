package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex
}

func (p *post) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock()
	p.views += 1
	p.mu.Unlock()
}

func mutexExample() {
	var wg sync.WaitGroup
	myPost := post{
		views: 0,
	}

	for range 100{
		wg.Add(1)
		go myPost.increment(&wg)
	}

	wg.Wait()
	// myPost.increment()
	fmt.Println(myPost.views)
}