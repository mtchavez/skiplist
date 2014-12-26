package skiplist

import (
	"fmt"
	"testing"
)

func BenchmarkInsert_1000(b *testing.B)    { benchInsert(b, 1000) }
func BenchmarkInsert_10000(b *testing.B)   { benchInsert(b, 10000) }
func BenchmarkInsert_100000(b *testing.B)  { benchInsert(b, 100000) }
func BenchmarkInsert_1000000(b *testing.B) { benchInsert(b, 1000000) }

func benchInsert(b *testing.B, total int) {
	for i := 0; i < b.N; i++ {
		l := NewList()
		for i := 0; i < total; i++ {
			l.Insert(i, []byte(fmt.Sprintf("Skiplist node insert %d", i)))
		}
	}
}

func BenchmarkParallelInsert(b *testing.B) {
	l := NewList()
	i := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i++
			l.Insert(i, []byte(fmt.Sprintf("Skiplist node insert %d", i)))
		}
	})
}

func BenchmarkDelete_1000(b *testing.B)    { benchDelete(b, 1000) }
func BenchmarkDelete_10000(b *testing.B)   { benchDelete(b, 10000) }
func BenchmarkDelete_100000(b *testing.B)  { benchDelete(b, 100000) }
func BenchmarkDelete_1000000(b *testing.B) { benchDelete(b, 1000000) }

func benchDelete(b *testing.B, total int) {
	l := NewList()
	for i := 0; i < total; i++ {
		l.Insert(i, []byte(fmt.Sprintf("Skiplist node insert %d", i)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < total; i++ {
			l.Delete(i)
		}
	}
}

func BenchmarkParallelDelete(b *testing.B) {
	l := NewList()
	for i := 0; i < 1000000; i++ {
		l.Insert(i, []byte(fmt.Sprintf("Skiplist node insert %d", i)))
	}
	b.ResetTimer()
	i := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i++
			l.Delete(i)
		}
	})
}
