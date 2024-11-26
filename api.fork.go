package gokzg4844

import (
	"slices"

	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
)

// RootsOfUnity returns the roots of unity used for the evaluation points of the
// KZG polynomial.
func (c *Context) RootsOfUnity() []fr.Element {
	return slices.Clone(c.domain.Roots)
}
