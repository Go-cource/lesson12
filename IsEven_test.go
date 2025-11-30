package main

import "testing"

func BenchmarkInsertXInMap100(b *testing.B) {
	for i := 0; i < b.N; i++{
		InsertXInMap(100)
	}
}
func BenchmarkInsertXInMap1000(b *testing.B) {
	for i := 0; i < b.N; i++{
		InsertXInMap(1000)
	}
}

func BenchmarkInsertXInMap100000(b *testing.B) {
	for i := 0; i < b.N; i++{
		InsertXInMap(100000)
	}
}
