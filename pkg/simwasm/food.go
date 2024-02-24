package simwasm

import "simcore"

type Food struct {
	X float32
	Y float32
}

// FromCoreFood implements conversion of *simcore.Food to Food.
func (Food) FromCoreFood(food *simcore.Food) Food {
	return Food{
		food.GetPosition().X,
		food.GetPosition().Y,
	}
}
