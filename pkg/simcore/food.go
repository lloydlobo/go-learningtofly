package simcore

import (
	"math/rand"

	"simcore/internal/nalgebra"
)

type Food struct {
	Position nalgebra.Point2
}

func (f *Food) Random(rng *rand.Rand) Food {
	return Food{nalgebra.Point2{}.Random(rng)}
}

// GetPosition implements a getter to access Position from Food object's state.
func (f *Food) GetPosition() nalgebra.Point2 {
	return f.Position
}
