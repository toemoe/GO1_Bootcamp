package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Worker struct {
	id    int
	sleep int
}

type WorkerManager struct {
	results []Worker
	mu      sync.Mutex
	wg      sync.WaitGroup
}

func (wm *WorkerManager) sleepGoroutine(id, m int) {
	defer wm.wg.Done()
	sleepTime := rand.Intn(m) + 1
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	wm.mu.Lock()
	wm.results = append(wm.results, Worker{id: id, sleep: sleepTime})
	wm.mu.Unlock()
}

func main() {
	n := flag.Int("N", 3, "Number of goroutines")
	m := flag.Int("M", 10, "Max sleep time in ms")
	flag.Parse()
	if *n <= 0 || *m <= 0 {
		fmt.Println("error: N and M must be greater than 0")
		return
	}

	wm := &WorkerManager{
		results: make([]Worker, 0, *n),
	}

	for i := 0; i < *n; i++ {
		wm.wg.Add(1)
		go wm.sleepGoroutine(i, *m)
	}

	wm.wg.Wait()
	sort.Slice(wm.results, func(i, j int) bool {
		return wm.results[i].sleep > wm.results[j].sleep
	})

	for _, r := range wm.results {
		fmt.Printf("<%d, %dms>\n", r.id, r.sleep)
	}
}
