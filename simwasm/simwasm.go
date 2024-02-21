package simwasm

import "math/rand"

// FUTURE: use lib_simulation as sim;

// #[wasm_bindgen]
type Simulation struct {
	Rng rand.Rand
	// sim sim.Simulation
}
