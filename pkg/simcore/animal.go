package simcore

import (
	"math/rand"

	"simcore/internal/nalgebra"
)

type Animal struct {
	Position nalgebra.Point2
	Rotation nalgebra.Rotation2
	Speed    float32
}

func (a *Animal) Random(rng *rand.Rand) Animal {
	const speed = 0.002

	return Animal{
		nalgebra.Point2{}.Random(rng),
		nalgebra.NewRotation2(rng.Float32()),
		speed,
	}
}

// GetPosition implements getter to access Position from Animal object's state.
func (a *Animal) GetPosition() nalgebra.Point2 {
	return a.Position
}

// GetRotation implements getter to access Rotation from Animal object's state.
func (a *Animal) GetRotation() nalgebra.Rotation2 {
	return a.Rotation
}
