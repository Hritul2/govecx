package govecx

import "testing"

// Run with: go test -bench=. -benchmem -run=^$

func BenchmarkFlatSearch(b *testing.B) {
	b.Skip("TODO(step 6): N=10k, dim=384, k=10")
}

func BenchmarkHNSWSearch(b *testing.B) {
	b.Skip("TODO(step 6): N=10k/100k, dim=384, k=10")
}

func BenchmarkHNSWAdd(b *testing.B) {
	b.Skip("TODO(step 6): build time per vector")
}
