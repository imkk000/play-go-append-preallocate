package poc

import "testing"

const (
	maxSize = int(1e7)
)

func Benchmark(b *testing.B) {
	b.Run("Append", func(b *testing.B) {
		b.StartTimer()
		var v []int
		for i := 0; i < maxSize; i++ {
			v = append(v, i)
		}
		b.StopTimer()

		l := len(v)
		if l != maxSize {
			b.Errorf("want = %d got %d", l, maxSize)
		}
		if v[l-1] != maxSize-1 {
			b.Errorf("want = %d got %d", maxSize, v[l-1])
		}
	})

	b.Run("Preallocate", func(b *testing.B) {
		b.StartTimer()
		v := make([]int, maxSize)
		for i := 0; i < maxSize; i++ {
			v[i] = i
		}
		b.StopTimer()

		l := len(v)
		if l != maxSize {
			b.Errorf("want = %d got %d", l, maxSize)
		}
		if v[l-1] != maxSize-1 {
			b.Errorf("want = %d got %d", maxSize, v[l-1])
		}
	})
}
