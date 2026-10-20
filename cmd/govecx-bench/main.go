// Command govecx-bench measures QPS and recall@k for Flat vs HNSW.
//
//	go run ./cmd/govecx-bench -n 100000 -dim 384 -k 10
package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 10_000, "number of vectors")
	dim := flag.Int("dim", 384, "vector dimension")
	k := flag.Int("k", 10, "neighbors per query")
	queries := flag.Int("queries", 1_000, "number of queries")
	flag.Parse()

	// TODO(step 6):
	//   1. generate (or load from testdata/) n random vectors + queries
	//   2. build Flat and HNSW, time the build
	//   3. run queries on both, compute QPS
	//   4. recall@k = |HNSW ∩ Flat| / k, averaged over queries
	//   5. print a markdown table row: N | dim | build | QPS | recall@k
	fmt.Printf("TODO: bench n=%d dim=%d k=%d queries=%d\n", *n, *dim, *k, *queries)
}
