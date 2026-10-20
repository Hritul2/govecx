package govecx

// Metric is the similarity function used to compare vectors.
type Metric int

const (
	Cosine Metric = iota
	DotProduct
	Euclidean
)

func (m Metric) String() string {
	switch m {
	case Cosine:
		return "cosine"
	case DotProduct:
		return "dot"
	case Euclidean:
		return "euclidean"
	default:
		return "unknown"
	}
}

// TODO(step 3): add a helper that maps a Metric to a distance func from
// internal/distance, e.g. func (m Metric) distFunc() func(a, b []float32) float32
