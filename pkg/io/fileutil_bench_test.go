package io_test

import (
	"path/filepath"
	"testing"

	optio "github.com/getoptimum/optimum-common/pkg/io"
)

var benchFileData []byte

func benchmarkFilePayload(size int) []byte {
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(i)
	}
	return data
}

func BenchmarkAtomicallySaveToFile(b *testing.B) {
	for _, tc := range []struct {
		name string
		size int
	}{
		{name: "1KiB", size: 1024},
		{name: "64KiB", size: 64 * 1024},
	} {
		b.Run(tc.name, func(b *testing.B) {
			payload := benchmarkFilePayload(tc.size)
			target := filepath.Join(b.TempDir(), "payload.bin")

			b.SetBytes(int64(len(payload)))
			b.ReportAllocs()
			b.ResetTimer()

			for range b.N {
				if err := optio.AtomicallySaveToFile(target, payload); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkLoadFromFile(b *testing.B) {
	for _, tc := range []struct {
		name string
		size int
	}{
		{name: "1KiB", size: 1024},
		{name: "64KiB", size: 64 * 1024},
	} {
		b.Run(tc.name, func(b *testing.B) {
			payload := benchmarkFilePayload(tc.size)
			target := filepath.Join(b.TempDir(), "payload.bin")
			if err := optio.AtomicallySaveToFile(target, payload); err != nil {
				b.Fatal(err)
			}

			b.SetBytes(int64(len(payload)))
			b.ReportAllocs()
			b.ResetTimer()

			for range b.N {
				var err error
				benchFileData, err = optio.LoadFromFile(target)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
