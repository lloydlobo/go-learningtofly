package main

import "simcore"

type World struct {
	Animals []Animal
	Foods   []Food
}

// FromCoreWorld implements conversion of *simcore.World to World.
func (World) FromCoreWorld(world *simcore.World) World {
	animals := make([]Animal, 0)
	foods := make([]Food, 0)

	for _, a := range world.Animals {
		animals = append(animals, Animal{}.FromCoreAnimal(&a))
	}
	for _, f := range world.Foods {
		foods = append(foods, Food{}.FromCoreFood(&f))
	}

	return World{
		Animals: animals,
		Foods:   foods,
	}
}
