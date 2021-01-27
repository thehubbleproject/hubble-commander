package simulator

import (
	"time"
)

func benchmark(n int64, f func(int64)) int64 {
	start := time.Now()
	for i := int64(0); i < n; i++ {
		f(i)
	}
	delta := time.Since(start)
	return int64(delta) / n
}

func benchmarkPerSecond(n int64, f func(int64)) int {
	dt := float64(benchmark(n, f))
	return int(1e9 / dt)
}
