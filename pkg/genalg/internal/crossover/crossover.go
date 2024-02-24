package crossover

import (
	"math/rand"

	"genalg/internal/chromosome"

	_ "genalg/internal/crossover/uniform"
)

type CrossoverMethod interface {
	Crossover(rng *rand.Rand, parentA, parentB chromosome.Chromosome) chromosome.Chromosome
}
