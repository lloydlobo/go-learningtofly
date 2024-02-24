package simwasm

import "simcore"

type Animal struct {
	X        float32
	Y        float32
	Rotation float32
}

// FromCoreAnimal implements conversion of *simcore.Animal to Animal.
func (Animal) FromCoreAnimal(animal *simcore.Animal) Animal {
	return Animal{
		animal.GetPosition().X,
		animal.GetPosition().Y,
		animal.GetRotation().Angle(),
	}
}

// ^ This model is smaller than `lib_simulation::Animal` (`simcore.Animal`) -
// | that's because a bird's position is all we need on the JavaScript's
// | side at the moment; there's no need to map rest of the fields.
