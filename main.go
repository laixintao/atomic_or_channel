package main

import (
	"fmt"
	"time"
)

var LOOP = 1000000
var CLIENTS = 100

func main() {
	start := time.Now()
	result := ChannelAdd()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("ChannelAdd() Done! count=%v took %vs\n", result, elapsed)

	start = time.Now()
	result2 := AtomicAdd()
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("AtomicAdd() Done! count=%v took %vs\n", result2, elapsed)
}
