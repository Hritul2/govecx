// Package distance holds the vector math hot loops.
//
// Convention to decide on (step 3): return a *distance* (lower = closer)
// so every metric can share the same heap logic, and convert to a
// user-facing score at the API boundary.
package distance

// Func compares two equal-length vectors.
type Func func(a, b []float32) float32

// Dot returns the dot product of a and b.
func Dot(a, b []float32) float32 {
	// TODO(step 3)
	panic("TODO")
}

// SquaredL2 returns the squared Euclidean distance (skip the sqrt — it doesn't change ordering).
func SquaredL2(a, b []float32) float32 {
	// TODO(step 3)
	panic("TODO")
}

// CosineDistance returns 1 - cosine similarity.
// Tip: normalize vectors on Add, then cosine == dot and this gets much cheaper.
func CosineDistance(a, b []float32) float32 {
	// TODO(step 3)
	panic("TODO")
}

// Normalize scales v to unit length in place.
func Normalize(v []float32) {
	// TODO(step 3)
	panic("TODO")
}
