// Package testutil provides utilities for testing floating-point numbers.
package testutil

import (
	"math"
)

var (
	// Float32Epsilon is the smallest positive value representable by a float32,
	// used as tolerance for comparisons. Its value is approximately 1.1920929e-07.
	Float32Epsilon = math.Nextafter32(1, 2) - 1
)

// ApproxRelativeEq checks if two float32 numbers are approximately equal
// within a specified tolerance. It returns true if the absolute difference
// between them is less than the product of the tolerance and the maximum magnitude.
//
// See https://floating-point-gui.de/
func ApproxRelativeEq(left, right, tolerance float32) bool {
	diff := math.Abs(float64(left - right))
	max := math.Max(math.Abs(float64(left)), math.Abs(float64(right)))

	return diff <= float64(tolerance*float32(max))
}
