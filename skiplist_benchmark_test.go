package skiplist

import (
	"fmt"
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	l := New()
	for i := 0; i < b.N; i++ {
		l.Insert(i, []byte(fmt.Sprintf("Skiplist node insert %i", i)))
	}
}

func BenchmarkDelete(b *testing.B) {
	l := New()
	for i := 0; i < 1000000; i++ {
		l.Insert(i, []byte(fmt.Sprintf("Skiplist node insert %i", i)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Delete(i)
	}
}
