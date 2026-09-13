package syncx_test

import (
	"strconv"
	"testing"

	"github.com/getoptimum/optimum-common/pkg/syncx"
)

var (
	benchRWMapInt int
	benchRWMapAll map[string]int
)

func newBenchRWMap(size int) *syncx.RWMap[string, int] {
	m := syncx.NewRWMap[string, int]()
	for i := range size {
		m.Store("key-"+strconv.Itoa(i), i)
	}
	return m
}

func BenchmarkRWMapLoad(b *testing.B) {
	const size = 10_000
	m := newBenchRWMap(size)

	b.ReportAllocs()
	b.ResetTimer()
	for i := range b.N {
		benchRWMapInt, _ = m.Load("key-" + strconv.Itoa(i%size))
	}
}

func BenchmarkRWMapLoadParallel(b *testing.B) {
	const size = 10_000
	m := newBenchRWMap(size)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		value := 0
		for pb.Next() {
			value, _ = m.Load("key-" + strconv.Itoa(i%size))
			i++
		}
		_ = value
	})
}

func BenchmarkRWMapStore(b *testing.B) {
	const size = 10_000
	m := syncx.NewRWMap[string, int]()

	b.ReportAllocs()
	b.ResetTimer()
	for i := range b.N {
		m.Store("key-"+strconv.Itoa(i%size), i)
	}
}

func BenchmarkRWMapLoadAll(b *testing.B) {
	const size = 10_000
	m := newBenchRWMap(size)

	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		benchRWMapAll = m.LoadAll()
	}
}

func BenchmarkRWMapMixed80R20W(b *testing.B) {
	const size = 10_000
	m := newBenchRWMap(size)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		value := 0
		for pb.Next() {
			key := "key-" + strconv.Itoa(i%size)
			if i%5 == 0 {
				m.Store(key, i)
			} else {
				value, _ = m.Load(key)
			}
			i++
		}
		_ = value
	})
}
