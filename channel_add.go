package main

import (
	"sync"
	"sync/atomic"
)

func ChannelAdd() uint64 {
	var wg sync.WaitGroup
	var count uint64

	count = 0

	numCh := make(chan *uint64, 10240)
	// start worker
	go func() {
		for value := range numCh {
			atomic.AddUint64(&count, *value)
		}
	}()

	one := uint64(1)
	for i := 1; i <= CLIENTS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for l := 0; l < LOOP; l++ {
				numCh <- &one
			}
		}()
	}

	wg.Wait()
	close(numCh)
	return count

}
