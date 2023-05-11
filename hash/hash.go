// Package hash implements a deep hash mechanism so that arbitrary types can be
// used as a key in hash maps
package hash

import "moqueries.org/deephash"

// Hash stores the hash of a source object
type Hash uint64

// DeepHash walks the src parameter and produces a hash
func DeepHash(src interface{}) Hash {
	return Hash(deephash.Hash(src))
}

// DeepDiff returns a list of differences between lSrc and rSrc
func DeepDiff(lSrc, rSrc interface{}) []string {
	return deephash.Diff("", lSrc, rSrc)
}
