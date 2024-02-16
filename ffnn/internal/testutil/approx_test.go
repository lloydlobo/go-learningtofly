package testutil

import (
	"testing"
)

func TestApproxRelativeEq(t *testing.T) {
	tests := []struct {
		left      float32
		right     float32
		tolerance float32
		want      bool
	}{
		{0.1, 0.100001, 0.00001, true},                     // Within tolerance
		{0.1, 0.10001, 0.000001, false},                    // Outside tolerance
		{0.00001, 0.00002, 0.000001, false},                // Outside tolerance
		{1000, 1001, 0.0001, false},                        // Outside tolerance
		{0.45, (0.15 + 0.15 + 0.15), Float32Epsilon, true}, // Within tolerance
		{0.0, 0.0, 0.0, true},                              // Zero values
		{0.0, (-1.0 * 0.0), 0.0, true},                     // Positive and negative zero
		{-0.00001, -0.00001, 0.000001, true},               // Negative values within tolerance
		{0.0, 1.0, 0.000001, false},                        // Values with large difference
		{1.0e-10, 2.0e-10, 1.0e-11, false},                 // Very small values outside tolerance
	}

	for _, tt := range tests {
		got := ApproxRelativeEq(tt.left, tt.right, tt.tolerance)

		if got != tt.want {
			t.Errorf("ApproxRelativeEq(%f, %f, %f) = %t; want %t", tt.left, tt.right, tt.tolerance, got, tt.want)
		}
	}
}
