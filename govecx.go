package govecx

// Vector is a single embedding.
type Vector = []float32

// Result is one search hit.
type Result struct {
	ID    string
	Score float32 // higher = more similar (decide and document this convention)
}

// Index is the public contract every index type implements.
type Index interface {
	// Add inserts a vector under an external ID.
	Add(id string, v Vector) error

	// Search returns up to k results, best first.
	Search(q Vector, k int) ([]Result, error)

	// Len returns the number of stored vectors.
	Len() int
}

// Algorithm selects the index implementation.
type Algorithm int

const (
	HNSW Algorithm = iota
	Flat
)

// New builds an index from options.
func New(opts ...Option) (Index, error) {
	cfg := defaultConfig()
	for _, o := range opts {
		o(&cfg)
	}
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	// TODO(step 4/5): return newFlat(cfg) or newHNSW(cfg) based on cfg.algorithm.
	return nil, ErrNotImplemented
}
