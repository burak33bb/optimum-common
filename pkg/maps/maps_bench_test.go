package maps_test

import (
	"strconv"
	"testing"

	optmaps "github.com/getoptimum/optimum-common/pkg/maps"
)

var (
	benchMapKeys   []string
	benchMapValues []int
)

func benchmarkStdMap(size int) map[string]int {
	m := make(map[string]int, size)
	for i := range size {
		m["key-"+strconv.Itoa(i)] = i
	}
	return m
}

func BenchmarkMapKeys(b *testing.B) {
	for _, tc := range []struct {
		name string
		size int
	}{
		{name: "10", size: 10},
		{name: "1K", size: 1000},
		{name: "10K", size: 10_000},
	} {
		b.Run(tc.name, func(b *testing.B) {
			m := benchmarkStdMap(tc.size)
			b.ReportAllocs()
			b.ResetTimer()

			for range b.N {
				benchMapKeys = optmaps.MapKeys(m)
			}
		})
	}
}

func BenchmarkMapValues(b *testing.B) {
	for _, tc := range []struct {
		name string
		size int
	}{
		{name: "10", size: 10},
		{name: "1K", size: 1000},
		{name: "10K", size: 10_000},
	} {
		b.Run(tc.name, func(b *testing.B) {
			m := benchmarkStdMap(tc.size)
			b.ReportAllocs()
			b.ResetTimer()

			for range b.N {
				benchMapValues = optmaps.MapValues(m)
			}
		})
	}
}
