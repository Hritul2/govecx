// Package govecx is an in-memory nearest-neighbor search library for Go.
//
// Think of it as "the map of vectors": add embeddings with an ID, then ask
// for the k closest ones to a query vector. No external services, pure Go.
//
//	idx, err := govecx.New(
//		govecx.WithDimension(768),
//		govecx.WithMetric(govecx.Cosine),
//	)
//	err = idx.Add("doc-123", embedding)
//	results, err := idx.Search(query, 10)
//	for _, r := range results {
//		fmt.Println(r.ID, r.Score)
//	}
package govecx
