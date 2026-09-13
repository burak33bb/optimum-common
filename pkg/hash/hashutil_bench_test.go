package hash_test

import (
	"crypto/sha256"
	"testing"

	"github.com/getoptimum/optimum-common/pkg/hash"
)

var (
	benchHashString string
	benchHashArray  [sha256.Size]byte
	benchHashUint64 uint64
)

var benchHashSizes = []struct {
	name string
	size int
}{
	{name: "64B", size: 64},
	{name: "1KiB", size: 1024},
	{name: "64KiB", size: 64 * 1024},
}

func benchmarkPayload(size int) []byte {
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(i)
	}
	return data
}

func BenchmarkSHA256Sizes(b *testing.B) {
	for _, tc := range benchHashSizes {
		b.Run(tc.name, func(b *testing.B) {
			data := benchmarkPayload(tc.size)
			b.SetBytes(int64(len(data)))
			b.ReportAllocs()
			b.ResetTimer()

			for range b.N {
				benchHashString = hash.SHA256(data)
			}
		})
	}
}

func BenchmarkSHA512Sizes(b *testing.B) {
	for _, tc := range benchHashSizes {
		b.Run(tc.name, func(b *testing.B) {
			data := benchmarkPayload(tc.size)
			b.SetBytes(int64(len(data)))
			b.ReportAllocs()
			b.ResetTimer()

			for range b.N {
				benchHashString = hash.SHA512(data)
			}
		})
	}
}

func BenchmarkSHA256StringSizes(b *testing.B) {
	for _, tc := range benchHashSizes {
		b.Run(tc.name, func(b *testing.B) {
			data := benchmarkPayload(tc.size)
			b.SetBytes(int64(len(data)))
			b.ReportAllocs()
			b.ResetTimer()

			for range b.N {
				benchHashArray = hash.SHA256String(data)
			}
		})
	}
}

func BenchmarkBytesHashSizes(b *testing.B) {
	for _, tc := range []struct {
		name string
		size int
	}{
		{name: "Empty", size: 0},
		{name: "64B", size: 64},
		{name: "1KiB", size: 1024},
		{name: "64KiB", size: 64 * 1024},
	} {
		b.Run(tc.name, func(b *testing.B) {
			data := benchmarkPayload(tc.size)
			b.SetBytes(int64(len(data)))
			b.ReportAllocs()
			b.ResetTimer()

			for range b.N {
				benchHashString = hash.BytesHash(data)
			}
		})
	}
}

func BenchmarkXXHashSizes(b *testing.B) {
	for _, tc := range benchHashSizes {
		b.Run(tc.name, func(b *testing.B) {
			data := benchmarkPayload(tc.size)
			b.SetBytes(int64(len(data)))
			b.ReportAllocs()
			b.ResetTimer()

			for range b.N {
				benchHashUint64 = hash.XXHash(data)
			}
		})
	}
}
