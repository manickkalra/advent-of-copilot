package main

import (
	"testing"
)

// benchmark the main file
func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
