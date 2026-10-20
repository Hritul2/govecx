package govecx

import "testing"

func TestNewValidatesOptions(t *testing.T) {
	t.Skip("TODO(step 2): dimension <= 0, M < 2, efSearch < 1 should error")
}

func TestFlatAddSearch(t *testing.T) {
	t.Skip("TODO(step 4): add a few hand-made vectors, check exact nearest neighbors")
}

func TestHNSWRecall(t *testing.T) {
	t.Skip("TODO(step 6): 10k random vectors, compare HNSW vs Flat, require recall@10 >= 0.95")
}

func TestErrors(t *testing.T) {
	t.Skip("TODO: dimension mismatch, duplicate ID, k <= 0, search on empty index")
}
