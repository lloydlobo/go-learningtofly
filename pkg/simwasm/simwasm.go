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

// Note: Rest of the code in this directory is somewhat copy-pasted from
// lib-simulation(simcore), but with WebAssembly in mind...
