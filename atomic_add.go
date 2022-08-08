package main

import (
	"sync"
	"sync/atomic"
)

func AtomicAdd() uint64 {
	var wg sync.WaitGroup
	var count uint64

	count = 0

	for i := 1; i <= CLIENTS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for l := 0; l < LOOP; l++ {
				atomic.AddUint64(&count, 1)
			}
		}()
	}

	wg.Wait()
	return count
}
