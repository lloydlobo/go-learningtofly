package genalg

import (
	"math/rand"

	"genalg/internal/crossover"
	"genalg/internal/individual"
	"genalg/internal/selection"
)

func New[S SelectionMethod](selectionMethod S) GeneticAlgo[S] {
	return GeneticAlgo[S]{
		selectionMethod: selectionMethod,
		crossoverMethod: nil,
	}
}

type (
	SelectionMethod selection.SelectionMethod[individual.Individual]

	GeneticAlgo[S SelectionMethod] struct {
		// Note: selection algorithm generally remains identical for the whole simulation, so it'll be wiser to go with constructor
		selectionMethod S

		// Note: Crossover doesn't actually happen on an individual, but rather on their chromosomes
		crossoverMethod crossover.CrossoverMethod

		// mutationMethod MutationMethod[I]
	}
)

func (ga GeneticAlgo[S]) Evolve(
	rng *rand.Rand,
	population []individual.Individual,
	fitnessFn func(indiv individual.Individual) float32,
) []individual.Individual {

	populationCount := len(population)
	if populationCount == 0 {
		panic("expected population to not be empty")
	}

	output := make([]individual.Individual, populationCount)

	for i := range population {
		// 1. selection
		parentA := (*ga.selectionMethod.Select(rng, &population)).Chromosome()
		parentB := (*ga.selectionMethod.Select(rng, &population)).Chromosome()

		// 2. crossover
		child := ga.crossoverMethod.Crossover(rng, *parentA, *parentB)

		// 3. mutation
		// 	TODO
		// ga.mutationMethod.Mutate(rng, &child)

		var indiv individual.Individual
		output[i] = indiv.Create(child)
	}

	return output // , Statistic.New(population)
}
