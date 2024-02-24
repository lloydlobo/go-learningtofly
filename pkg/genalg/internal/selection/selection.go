// Any exported symbols (types, functions, variables) from roulette_wheel package
// will be accessible from the selection package without explicitly
// referencing roulette_wheel.
package selection

import (
	"math/rand"

	"genalg/internal/individual"

	// The underscore _ before the package path tells Go that you're importing the
	// package solely for its side effects, such as registering with the main package.
	// This ensures that the roulette_wheel package is imported and its init
	// functions are executed, but its symbols aren't directly accessible from
	// selection unless explicitly exported.
	_ "genalg/internal/selection/roulette_wheel"
)

type SelectionMethod[I individual.Individual] interface {
	Select(rng *rand.Rand, population *[]I) I
}
