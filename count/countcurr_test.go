package main

import "testing"

func BenchmarkDoMuxAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doMuxAdd()
	}
}

func BenchmarkDoAtoAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doAtoAdd()
	}
}
