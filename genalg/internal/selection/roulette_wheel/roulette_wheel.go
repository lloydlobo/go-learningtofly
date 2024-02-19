package roulettewheel

import (
	"math/rand"

	"genalg/internal/individual"
)

type RouletteWheelSelection struct{}

func (r *RouletteWheelSelection) Select(
	rng *rand.Rand,
	population *[]individual.Individual,
) *individual.Individual { // ? how to make Individual type generic?
	if len(*population) == 0 {
		panic("population is empty")
	}

	var totalFitness float32
	for _, indiv := range *population {
		totalFitness += indiv.Fitness()
	}

	// This is a naïve approach for demonstration purposes; a more
	// efficient implementation could invoke `rng` just once
	for {
		// Intn -> a non-negative pseudo-random number in the half-open interval [0,n).
		indiv := (*population)[rng.Intn(len(*population))] // It panics if n <= 0.
		indivShare := indiv.Fitness() / totalFitness

		if rng.Float32() < indivShare { // rng.Float32() -> half-open interval [0.0,1.0)
			return &indiv
		}
	}
}