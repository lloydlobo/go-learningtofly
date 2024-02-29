// wasm-bindgen doesn't currently support exporting vectors of custom types.
//
// Even if it did, it's essential to maintain separation of concerns.
// lib-simulation (simcore) should focus on simulating evolution, not
// integrating with WebAssembly.
//
// By keeping lib-simulation(simcore) frontend-agnostic, it's easier to create
// different frontends such as lib-simulation-bevy or lib-simulation-cli
// sharing the same simulation code.
//
// Example:
//
//	#[wasm_bindgen]
//	#[derive(Debug)]
//	struct World {
//	    // Ok:
//	    animals: Vec<Animal>,
//
//	    // Error:
//	    pub foods: Vec<Food>,
//	}

package simcore

import (
	"math/rand"
)

type World struct {
	Animals []Animal
	Foods   []Food
}

func (w *World) Random(rng *rand.Rand) World {
	const (
		animalCount = 40
		foodCount   = 60
	)

	animals := make([]Animal, animalCount)
	for i := 0; i < animalCount; i++ {
		animals[i] = (&Animal{}).Random(rng)
	}

	foods := make([]Food, foodCount)
	for i := 0; i < foodCount; i++ {
		foods[i] = (&Food{}).Random(rng)
	}

	// ^ Our algorithm allows for animals and foods to overlap, so
	// | it's hardly ideal - but good enough for our purposes.
	// | A more complex solution could be based off of e.g.
	// | Poisson disk sampling:
	// | https://en.wikipedia.org/wiki/Supersampling
	// ---

	return World{animals, foods}
}

// GetAnimals implements a getter to access Animals from World object's state.
func (w *World) GetAnimals() *[]Animal {
	return &w.Animals
}

// GetFoods implements a getter to access Foods from World object's state.
func (w *World) GetFoods() *[]Food {
	return &w.Foods
}
