package nalgebra

import "math"

type Rotation2 float32

func NewRotation2(angle float32) Rotation2 {
	// Normalize angle to be within [0, 2π)
	angle = float32(math.Mod(float64(angle), 2*math.Pi))
	if angle < 0.0 {
		angle += 2 * math.Pi
	}

	return Rotation2(angle)
}

func (r Rotation2) Angle() float32 {
	// Note: Is Rotation2 correctly implemented?
	// | Is Rotation2 a struct or just a type like enum?
	// |---v--------v
	return float32(r)
}
