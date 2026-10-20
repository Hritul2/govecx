# govecx — Build Steps

Goal: **pure Go, zero services, simple API, surprisingly fast** in-memory nearest-neighbor search.
Not a vector database. Just the index primitive.

Each step ends with a **✅ Done when** checklist. Don't move on until it passes.

---

## Project layout

```text
govecx/
├── doc.go                  package docs (shows on pkg.go.dev)
├── govecx.go               public types: Vector, Result, Index, New()
├── options.go              functional options + config + validate()
├── metric.go               Metric enum (Cosine / DotProduct / Euclidean)
├── errors.go               sentinel errors
├── flat.go                 brute-force index (baseline)
├── hnsw.go                 HNSW index
├── *_test.go               unit tests, benchmarks, Example()
├── internal/
│   ├── distance/           Dot, SquaredL2, Cosine, Normalize  (hot loops)
│   └── topk/               min/max heaps, keep-best-k
├── cmd/govecx-bench/       CLI: QPS + recall table
├── examples/basic/         smallest runnable program
├── testdata/               datasets (git-ignored)
└── .github/workflows/      CI
```

Why `internal/`? Anything there can't be imported by users, so you can change it freely without breaking anyone. Only the root package is your public API — keep it small.

Commands (see `Makefile`): `make test`, `make bench`, `make vet`, `make example`.

---

## Step 0 — Repo hygiene (15 min)

- [x] `git init`, first commit, create GitHub repo `Hritul2/govecx` (must match the `module` line in `go.mod`)
- [x] Pick a license (MIT or Apache-2.0) → add `LICENSE`, update README
- [x] Consider lowering `go 1.26.7` in `go.mod` to something like `go 1.23` — the `go` line is the *minimum* version users need, and a library should not force the newest toolchain
- [x] Optional: install `staticcheck` (`go install honnef.co/go/tools/cmd/staticcheck@latest`)

✅ Done when: `go build ./... && go vet ./... && go test ./...` passes (everything is skipped for now) and CI is green on GitHub.

---

## Step 1 — Lock the public API (1 hr)

The API is the hardest thing to change later. Decide now:

- [ ] **ID type.** Scaffold uses `string` (user-friendly). Alternative: `uint64` (faster, smaller). You can keep strings externally and map them to internal `uint32` indexes — that's what the scaffold does.
- [ ] **Score convention.** Users expect "higher = more similar". Internally use *distance* (lower = closer) everywhere. Convert at the boundary:
  - cosine: `score = 1 - dist`
  - dot: `score = dot`
  - euclidean: `score = -dist` (or `1/(1+dist)`) — pick one and document it
- [ ] **Does `Add` copy the slice?** (Yes — otherwise users mutating their slice corrupts your index.)
- [ ] **Does Cosine normalize on insert?** (Yes — then cosine = dot, much faster.)
- [ ] Write godoc comments on every exported symbol.

✅ Done when: you could write the README usage section from the godoc alone.

---

## Step 2 — Options & validation (30 min)

File: `options.go`

- [ ] Implement `config.validate()`: dimension > 0, M ≥ 2, efConstruction ≥ M, efSearch ≥ 1, metric known
- [ ] Wrap errors so callers can `errors.Is(err, ErrInvalidDimension)`: `fmt.Errorf("%w: got %d", ErrInvalidDimension, d)`
- [ ] Fill `TestNewValidatesOptions` (table-driven)

Learn: the *functional options pattern* (Dave Cheney, "Functional options for friendly APIs").

✅ Done when: bad configs return wrapped sentinel errors; tests pass.

---

## Step 3 — Distance functions & heaps (2–3 hr)

Files: `internal/distance/`, `internal/topk/`

Distance:
- [ ] `Dot`, `SquaredL2`, `CosineDistance`, `Normalize`
- [ ] Decide: what does Normalize do with a zero vector? (Return an error from `Add` is the cleanest.)
- [ ] Table-driven tests + `BenchmarkDot768`
- [ ] Try manual loop unrolling (4 or 8 accumulators) and compare benchmarks. Use `benchstat` to compare runs properly.
- [ ] Hook up `Metric → distance.Func` in `metric.go`

Heaps:
- [ ] `MinHeap` and `MaxHeap` of `Item{Index, Distance}` using `container/heap`
- [ ] Bounded top-k: push into a max-heap of size k; if full and new item is closer than the top, replace the top
- [ ] Test: 1000 random items → exactly the 10 smallest, sorted

Learn: why squared L2 is fine (sqrt is monotonic); floating-point accumulation error.

✅ Done when: distance funcs are tested and benchmarked; heap test passes.

---

## Step 4 — Flat (brute-force) index (1–2 hr)

File: `flat.go`

This is your **ground truth**. HNSW will be measured against it forever.

- [ ] Storage: one contiguous `[]float32`; vector `i` = `vectors[i*dim : (i+1)*dim]`
- [ ] `Add`: check dim, reject duplicate ID, copy, normalize if cosine, append
- [ ] `Search`: check dim + k, scan all, keep top-k with the max-heap, sort best-first, convert distance → score
- [ ] Edge cases: empty index, k > Len()
- [ ] Wire `New()` to return `newFlat(cfg)` when `WithAlgorithm(Flat)`
- [ ] Fill `TestFlatAddSearch`, `TestErrors`, the `Example()` in `example_test.go`
- [ ] `make example` should print results

✅ Done when: `go run ./examples/basic` prints correct neighbors with `WithAlgorithm(govecx.Flat)`.

**🏷 Milestone: you have a working library.** Tag `v0.0.1` if you want.

---

## Step 5 — HNSW (the main event, 1–2 days)

File: `hnsw.go`

### 5a. Read first
- Paper: Malkov & Yashunin, *Efficient and robust approximate nearest neighbor search using HNSW graphs* — https://arxiv.org/abs/1603.09320. Focus on **Algorithms 1–5**.
- Pinecone's HNSW explainer is a gentle intro to read before the paper.
- Mental model: a skip list, but for graphs. Top layers = few nodes, long jumps. Layer 0 = every node, short hops.

### 5b. Data structures (start simple!)
- [ ] First version: `levels []uint8` (top level of each node) and `neighbors [][][]uint32` (`neighbors[node][level]`). Easy to debug.
- [ ] Max neighbors: `M` for layers > 0, `2*M` for layer 0 (paper recommendation).
- [ ] `levelMult = 1 / ln(M)`

### 5c. Algorithms — implement in this order
1. [ ] `randomLevel()` → `floor(-ln(rand.Float64()) * levelMult)`. Use a seeded `*rand.Rand` (from `cfg.seed`) so tests are reproducible.
2. [ ] `searchLayer(q, entry, ef, level)` — **Algorithm 2**. Greedy best-first search with a candidates min-heap, results max-heap, and a visited set. This is 80% of HNSW.
3. [ ] `selectNeighbors` — start with **Algorithm 3** (simple: take the M closest). Upgrade to **Algorithm 4** (heuristic: keeps diverse neighbors) later — it noticeably improves recall.
4. [ ] `Add` — **Algorithm 1**:
   - first node becomes the entry point
   - greedy descend from top level down to `nodeLevel+1` with ef=1
   - from `min(nodeLevel, maxLevel)` down to 0: `searchLayer` with `efConstruction`, select M neighbors, connect **both directions**
   - if a neighbor now has too many links, shrink its list (re-run selectNeighbors on it)
   - if `nodeLevel > maxLevel`, update entry point
5. [ ] `Search` — **Algorithm 5**: greedy descend with ef=1 to layer 1, then `searchLayer` on layer 0 with `max(efSearch, k)`, return top k.

### 5d. Visited set
- [ ] Start with `map[uint32]struct{}`. Later replace with a reusable `[]uint32` "epoch" array (store a generation counter per node; bump it each search instead of clearing) — big speedup.

### 5e. Debugging tips
- Test on 2-D points you can reason about first.
- Write a `checkGraph()` test helper: no self-links, no duplicate links, no neighbor list over capacity, all node levels ≤ maxLevel.
- If recall is bad: check bidirectional links and the pruning step first.

✅ Done when: `TestHNSWRecall` passes (recall@10 ≥ 0.95 on 10k random 128-d vectors with defaults).

---

## Step 6 — Benchmarks & recall (half a day)

Files: `benchmark_test.go`, `cmd/govecx-bench/main.go`

- [ ] Go benchmarks: `BenchmarkFlatSearch`, `BenchmarkHNSWSearch`, `BenchmarkHNSWAdd`. Build the index *outside* the timed loop (`b.ResetTimer()`); use `b.ReportMetric` for QPS.
- [ ] `govecx-bench` CLI: random data first, then real data from `testdata/` (SIFT1M, GloVe — see `testdata/README.md`). Random data is the *hardest* case for ANN; real embeddings cluster and score better.
- [ ] recall@k = average over queries of `|HNSW results ∩ Flat results| / k`
- [ ] Sweep `efSearch` (10, 20, 50, 100, 200) → print QPS vs recall. This tradeoff curve is the most interesting thing you'll publish.
- [ ] Fill the README benchmark table (include CPU model + Go version)
- [ ] Profile: `go test -bench=HNSWSearch -cpuprofile cpu.out` then `go tool pprof -http=: cpu.out`

✅ Done when: README has a real table and one efSearch/recall curve.

---

## Step 7 — Release v0.1.0

- [ ] README: install, usage, benchmarks, limitations ("not safe for concurrent writes" etc.)
- [ ] Godoc on everything exported; `Example()` passes
- [ ] `go vet`, `staticcheck`, `go test -race` clean
- [ ] `git tag v0.1.0 && git push --tags`
- [ ] Trigger pkg.go.dev: `GOPROXY=proxy.golang.org go list -m github.com/Hritul2/govecx@v0.1.0`
- [ ] Try it from a fresh folder: `go mod init demo && go get github.com/Hritul2/govecx@v0.1.0`

Semver reminder: while on `v0.x` you may break the API. At `v1` you can't (without a `/v2` module path).

**v0.1 scope:** float32 · cosine · HNSW + Flat · Add · Search · top-k · configurable M / efConstruction / efSearch · benchmarks.

---

## v0.2 — Make it practical

- [ ] **Dot product & Euclidean** fully supported in HNSW (should be mostly free after step 3)
- [ ] **Concurrent Search**: `sync.RWMutex` — many readers, one writer. Make the visited-set pool-able (`sync.Pool`) so searches don't share state. Test with `-race`.
- [ ] **Delete**: tombstones (mark deleted, skip in results, keep traversing through them). Real removal + graph repair is hard — defer it. Optionally a `Compact()` that rebuilds.
- [ ] **Update**: delete + add (simplest correct approach)
- [ ] **Persistence**: `idx.Save(w io.Writer)` / `govecx.Load(r io.Reader)` (+ `SaveFile`/`LoadFile` helpers)
  - Format: magic bytes `"GVCX"` + format version + config + counts + raw vectors + IDs + graph
  - Use `encoding/binary` little-endian; wrap in `bufio`; add a checksum (`hash/crc32`)
  - Test: save → load → identical search results
  - Fuzz the loader (`go test -fuzz`) — corrupted files must return errors, never panic
- [ ] **Flatten the graph**: replace `[][][]uint32` with one contiguous `[]uint32` + offsets. Measure memory (`runtime.ReadMemStats`) and GC time before/after.

---

## v0.3 — Make it fast

- [ ] **SIMD distance**: Go assembly (avo generator) for amd64 AVX2 / arm64 NEON, with a pure-Go fallback behind build tags. Keep the pure-Go version as the test oracle.
- [ ] **Quantization**: int8 scalar quantization (4× less memory); search on quantized vectors, re-rank top candidates with float32
- [ ] **mmap persistence**: load the index without reading it into heap (`syscall.Mmap` / `golang.org/x/exp/mmap`) — design your file layout in v0.2 so vectors are page-aligned and contiguous
- [ ] **Filtering**: `SearchWithFilter(q, k, func(id string) bool)` — filter during traversal, not after (post-filtering returns too few results)
- [ ] **Batch search**: `SearchBatch(qs []Vector, k)` with a worker pool

---

## Reference reading

| Topic | Source |
|---|---|
| HNSW paper | arxiv.org/abs/1603.09320 |
| Reference C++ impl | github.com/nmslib/hnswlib (`hnswalg.h`) — very readable |
| Benchmarks / datasets | github.com/erikbern/ann-benchmarks |
| Functional options | Dave Cheney's blog |
| Go benchmarking | `go doc testing.B`, `golang.org/x/perf/cmd/benchstat` |
| Profiling | go.dev/blog/pprof |
| Go assembly / SIMD | github.com/mmcloughlin/avo |
