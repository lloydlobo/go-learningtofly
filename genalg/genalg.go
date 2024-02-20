package genalg

import (
	"math/rand"

	"genalg/internal/crossover"
	"genalg/internal/individual"
	"genalg/internal/mutation"
	"genalg/internal/selection"
)

type (
	// GeneticAlgorithm represents a genetic algorithm with selection, crossover, and mutation methods.
	GeneticAlgorithm[S SelectionMethod] struct {
		selectionMethod S                         // Selection algorithm generally remains identical for the whole simulation
		crossoverMethod crossover.CrossoverMethod // Crossover happens on chromosomes
		mutationMethod  mutation.MutationMethod
	}

	// SelectionMethod defines the type alias for selection methods.
	SelectionMethod selection.SelectionMethod[individual.Individual]
)

// New creates a new instance of GeneticAlgorithm.
func New[S SelectionMethod](selector S, crosser crossover.CrossoverMethod, mutator mutation.MutationMethod) GeneticAlgorithm[S] {
	return GeneticAlgorithm[S]{
		selectionMethod: selector,
		crossoverMethod: crosser,
		mutationMethod:  mutator,
	}
}

// Evolve performs the evolution process on the given population.
func (ga GeneticAlgorithm[S]) Evolve(rng *rand.Rand, population []individual.Individual, fitnessFn func(indiv individual.Individual) float32) []individual.Individual {
	populationCount := len(population)
	if populationCount == 0 {
		panic("expected population to not be empty")
	}

	output := make([]individual.Individual, populationCount)

	for i := range population {
		// 1. Selection - roulette wheel
		parentA := (*ga.selectionMethod.Select(rng, &population)).Chromosome()
		parentB := (*ga.selectionMethod.Select(rng, &population)).Chromosome()

		// 2. Crossover - uniform
		child := ga.crossoverMethod.Crossover(rng, *parentA, *parentB)

		// 3. Mutation - gaussian
		ga.mutationMethod.Mutate(rng, &child)

		// Create individual with mutated chromosome
		var indiv individual.Individual
		output[i] = indiv.Create(child)
	}

	return output // , Statistic.New(population)
}
