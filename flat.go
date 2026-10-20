package govecx

// flat is the brute-force index: compare the query against every vector.
// It is the correctness baseline that HNSW recall is measured against.
type flat struct {
	cfg config

	vectors []float32         // contiguous: vector i lives at [i*dim : (i+1)*dim]
	ids     []string          // internal index -> external ID
	idToIdx map[string]uint32 // external ID -> internal index
}

var _ Index = (*flat)(nil)

func newFlat(cfg config) *flat {
	// TODO(step 4)
	return &flat{cfg: cfg}
}

func (f *flat) Add(id string, v Vector) error {
	// TODO(step 4)
	return ErrNotImplemented
}

func (f *flat) Search(q Vector, k int) ([]Result, error) {
	// TODO(step 4)
	return nil, ErrNotImplemented
}

func (f *flat) Len() int { return len(f.ids) }
