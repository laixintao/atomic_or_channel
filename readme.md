Testing the performance of:

- multiple threads write into channels, and one goroutine do the work

vs

- just use `atomic` lib


## Run

```shell
$ go run .                                                  
ChannelAdd() Done! count=99991670 took 33.95414173ss         
AtomicAdd() Done! count=100000000 took 1.92393346ss
```

## Benchmark

```shell
$ go test -bench=. -count=10
goos: darwin                                  
goarch: amd64                                 
pkg: github.com/laixintao/atomic_or_channel
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkAtomic-12             1        1947990739 ns/op
BenchmarkAtomic-12             1        1957558508 ns/op
BenchmarkAtomic-12             1        1984851549 ns/op
BenchmarkAtomic-12             1        1957678020 ns/op
BenchmarkAtomic-12             1        1953746756 ns/op
BenchmarkAtomic-12             1        1957197155 ns/op
BenchmarkAtomic-12             1        1964173594 ns/op
BenchmarkAtomic-12             1        1962058130 ns/op
BenchmarkAtomic-12             1        2342306683 ns/op
BenchmarkAtomic-12             1        2313845157 ns/op
BenchmarkChannel-12            1        35211801199 ns/op
BenchmarkChannel-12            1        40597364557 ns/op
BenchmarkChannel-12            1        38452709531 ns/op
BenchmarkChannel-12            1        40201893971 ns/op
BenchmarkChannel-12            1        41802617846 ns/op
BenchmarkChannel-12            1        41463031707 ns/op
BenchmarkChannel-12            1        41985476702 ns/op
BenchmarkChannel-12            1        43106978329 ns/op
BenchmarkChannel-12            1        45582670783 ns/op
BenchmarkChannel-12            1        43751655673 ns/op
PASS                                          
ok      github.com/laixintao/atomic_or_channel  432.885s
```


Result: simple atomic is 15x faster than channels.
