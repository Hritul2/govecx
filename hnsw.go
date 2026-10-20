package govecx

// hnsw implements Hierarchical Navigable Small World graphs.
// Paper: Malkov & Yashunin, https://arxiv.org/abs/1603.09320
type hnsw struct {
	cfg config

	// Storage — keep it contiguous, avoid []*node.
	vectors []float32 // vector i at [i*dim : (i+1)*dim]
	ids     []string
	idToIdx map[string]uint32

	// Graph. Design this yourself (see docs/steps.md, step 5).
	// Start simple, e.g. levels []uint8 and neighbors [][][]uint32,
	// then flatten into one []uint32 once it works.

	entryPoint uint32
	maxLevel   int
	levelMult  float64 // 1 / ln(M)
}

var _ Index = (*hnsw)(nil)

func newHNSW(cfg config) *hnsw {
	// TODO(step 5)
	return &hnsw{cfg: cfg}
}

func (h *hnsw) Add(id string, v Vector) error {
	// TODO(step 5): Algorithm 1 (INSERT)
	return ErrNotImplemented
}

func (h *hnsw) Search(q Vector, k int) ([]Result, error) {
	// TODO(step 5): Algorithm 5 (K-NN-SEARCH)
	return nil, ErrNotImplemented
}

func (h *hnsw) Len() int { return len(h.ids) }

// Helpers you will need (signatures are suggestions — change freely):
//
//   func (h *hnsw) randomLevel() int
//   func (h *hnsw) searchLayer(q []float32, entry uint32, ef, level int) []candidate   // Algorithm 2
//   func (h *hnsw) selectNeighbors(q []float32, cands []candidate, m int) []uint32      // Algorithm 3 or 4
//   func (h *hnsw) vector(i uint32) []float32
