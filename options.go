package govecx

// Option configures an index (functional options pattern).
type Option func(*config)

type config struct {
	dimension int
	metric    Metric
	algorithm Algorithm

	// HNSW parameters
	m              int // max neighbors per node per layer
	efConstruction int // candidate list size while building
	efSearch       int // candidate list size while searching
	seed           int64
}

func defaultConfig() config {
	return config{
		metric:         Cosine,
		algorithm:      HNSW,
		m:              16,
		efConstruction: 200,
		efSearch:       50,
		seed:           42,
	}
}

func (c config) validate() error {
	// TODO(step 2): return ErrInvalidDimension / ErrInvalidOption for bad values.
	return nil
}

func WithDimension(d int) Option       { return func(c *config) { c.dimension = d } }
func WithMetric(m Metric) Option       { return func(c *config) { c.metric = m } }
func WithAlgorithm(a Algorithm) Option { return func(c *config) { c.algorithm = a } }
func WithM(m int) Option               { return func(c *config) { c.m = m } }
func WithEfConstruction(ef int) Option { return func(c *config) { c.efConstruction = ef } }
func WithEfSearch(ef int) Option       { return func(c *config) { c.efSearch = ef } }
func WithSeed(s int64) Option          { return func(c *config) { c.seed = s } }
