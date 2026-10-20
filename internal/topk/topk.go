// Package topk keeps the k best candidates seen so far.
//
// You need two heaps for HNSW:
//   - a min-heap of candidates to explore (closest first)
//   - a max-heap of current results, capped at size ef/k (farthest on top, so you can evict it)
//
// Use container/heap first; write a hand-rolled typed heap later if profiling says so.
package topk

// Item is a candidate: internal node index + its distance to the query.
type Item struct {
	Index    uint32
	Distance float32
}

// TODO(step 3): implement MinHeap and MaxHeap (container/heap.Interface),
// plus a bounded "keep best k" helper used by the flat index.
