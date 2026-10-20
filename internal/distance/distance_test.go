package distance

import "testing"

func TestDot(t *testing.T) {
	t.Skip("TODO(step 3): table-driven tests — orthogonal, identical, opposite vectors")
}

func TestSquaredL2(t *testing.T) {
	t.Skip("TODO(step 3)")
}

func TestCosineDistance(t *testing.T) {
	t.Skip("TODO(step 3): also test zero vectors — what should happen?")
}

func BenchmarkDot768(b *testing.B) {
	b.Skip("TODO(step 3): benchmark at dim 384/768/1536")
}
