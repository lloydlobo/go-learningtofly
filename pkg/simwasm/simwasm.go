package simwasm

import (
	"math/rand"

	"simcore"
)

// #//[wasm_bindgen]
type Simulation struct {
	Rng *rand.Rand
	Sim simcore.Simulation
}

// | #[wasm_bindgen]
// | impl Simulation {
// v ...

// #[wasm_bindgen(constructor)]
func New() *Simulation {
	// Original rand from Rust compile error:
	//  - rand internally depends on getrandom, which does support WebAssembly
	//    inside a web browser, but only when asked explicitly:
	rng := rand.New(rand.NewSource(0))
	sim := (&simcore.Simulation{}).Random(rng)

	return &Simulation{rng, sim}
}

// World implements returns converted Sim's simcore.World to World.
func (s *Simulation) World() World {
	return World{}.FromCoreWorld(&s.Sim.World)
}

// Step performs a single step - a single second, of simulation.
func (s *Simulation) Step() {
	s.Sim.Step(s.Rng)
}

// func (s Simulation) wasm_bindgen() {
// 	// wasm_bindgen implementation for Simulation goes here
// }
// ^
// | This function is not directly equivalent to the wasm_bindgen macro in Rust.
// | Instead, it serves as a placeholder for the wasm_bindgen implementation in Go.

// ^ ...
// | } // end of impl Simulation

// | Note: code below is kinda copy-pasted from lib-simulation(simcore),
// | but with WebAssembly in mind...
// v

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
