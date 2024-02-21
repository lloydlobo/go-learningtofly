package simwasm

import (
	"math/rand"

	"simcore"
)

// #[wasm_bindgen]
type Simulation struct {
	Rng *rand.Rand
	Sim simcore.Simulation
}

// #[wasm_bindgen]
// impl Simulation {
// ...

// #[wasm_bindgen(constructor)]
func New() *Simulation {
	// Original rand from Rust compile error:
	//  - rand internally depends on getrandom, which does support WebAssembly
	//    inside a web browser, but only when asked explicitly:
	rng := rand.New(rand.NewSource(0))
	sim := (&simcore.Simulation{}).Random(rng)

	return &Simulation{rng, sim}
}

// ...
// } // end of impl Simulation
