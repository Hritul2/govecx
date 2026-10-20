package govecx

import "errors"

var (
	ErrNotImplemented    = errors.New("govecx: not implemented")
	ErrInvalidDimension  = errors.New("govecx: invalid dimension")
	ErrDimensionMismatch = errors.New("govecx: vector dimension mismatch")
	ErrDuplicateID       = errors.New("govecx: duplicate id")
	ErrInvalidK          = errors.New("govecx: k must be > 0")
	ErrInvalidOption     = errors.New("govecx: invalid option")
)
