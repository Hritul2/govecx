# govecx

> The `map` of vectors for Go — fast in-memory nearest-neighbor search with a Go-native API.

Pure Go. Zero external services. Not a vector database — just the index.

> ⚠️ Work in progress. See [docs/steps.md](docs/steps.md) for the roadmap.

## Install

```sh
go get github.com/Hritul2/govecx
```

## Usage

```go
idx, err := govecx.New(
    govecx.WithDimension(768),
    govecx.WithMetric(govecx.Cosine),
)

err = idx.Add("doc-123", embedding)

results, err := idx.Search(query, 10)
for _, r := range results {
    fmt.Println(r.ID, r.Score)
}
```

## Benchmarks

| N    | dim | index | QPS | recall@10 |
|------|-----|-------|-----|-----------|
| 10K  | 384 | Flat  | –   | 1.00      |
| 10K  | 384 | HNSW  | –   | –         |
| 100K | 384 | HNSW  | –   | –         |
| 1M   | 384 | HNSW  | –   | –         |

## License

MIT — see [LICENSE](LICENSE).
