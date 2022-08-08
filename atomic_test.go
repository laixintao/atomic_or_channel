package main

import (
	"testing"
)

func BenchmarkAtomic(b *testing.B) {
	AtomicAdd()
}
