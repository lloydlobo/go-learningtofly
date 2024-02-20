package mutation

import (
	"math/rand"

	"genalg/internal/chromosome"
)

type MutationMethod interface {
	Mutate(rng *rand.Rand, child *chromosome.Chromosome)
}
