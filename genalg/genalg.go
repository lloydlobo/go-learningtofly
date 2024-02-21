package genalg

import (
	"math/rand"

	"genalg/internal/crossover"
	"genalg/internal/individual"
	"genalg/internal/mutation"
	"genalg/internal/selection"
	roulettewheel "genalg/internal/selection/roulette_wheel"
)

type (
	// GeneticAlgorithm represents a genetic algorithm with selection, crossover, and mutation methods.
	GeneticAlgorithm struct {
		selectionMethod selection.SelectionMethod[individual.Individual] // Selection algorithm generally remains identical for the whole simulation
		crossoverMethod crossover.CrossoverMethod                        // Crossover happens on chromosomes
		mutationMethod  mutation.MutationMethod
	}

	// SelectionMethod defines the type alias for selection methods.
	SelectionMethod selection.SelectionMethod[individual.Individual]
)

// New creates a new instance of GeneticAlgorithm.
func New(
	selector selection.SelectionMethod[individual.Individual],
	crosser crossover.CrossoverMethod,
	mutator mutation.MutationMethod,
) GeneticAlgorithm {
	return GeneticAlgorithm{
		selectionMethod: selector,
		crossoverMethod: crosser,
		mutationMethod:  mutator,
	}
}

// Evolve performs the evolution process on the given population.
func (ga GeneticAlgorithm) Evolve(
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
		// 1. Selection - roulette wheel: uses roulettewheel.RouletteWheel.Select()
		selectorA := roulettewheel.RouletteWheelSelection{}
		selectorB := roulettewheel.RouletteWheelSelection{}
		selectedA := selectorA.Select(rng, &population)
		selectedB := selectorB.Select(rng, &population)
		parentA := selectedA.Chromosome() // parentA := ga.selectionMethod.Select(rng, &population).Chromosome()
		parentB := selectedB.Chromosome() // parentB := ga.selectionMethod.Select(rng, &population).Chromosome()

		// 2. Crossover - uniform
		child := ga.crossoverMethod.Crossover(rng, *parentA, *parentB)

		// 3. Mutation - gaussian
		ga.mutationMethod.Mutate(rng, &child)

		// Note: solving this error. supressed by using TestIndividual mock struct
		// 		panic: runtime error: invalid memory address or nil pointer dereference [recovered]
		// 		        panic: runtime error: invalid memory address or nil pointer dereference
		/*
			next: [<nil> <nil> <nil> <nil>]
			tt.args.population: [<nil> <nil> <nil> <nil>]
			tt.args.population: [<nil> <nil> <nil> <nil>]
			--- FAIL: TestGeneticAlgorithm_Evolve/#00 (0.00s)
			--- FAIL: TestGeneticAlgorithm_Evolve (0.00s)
			panic: runtime error: invalid memory address or nil pointer dereference [recovered]
			        panic: runtime error: invalid memory address or nil pointer dereference
			[signal 0xc0000005 code=0x0 addr=0x28 pc=0xed2d02]
		*/

		// Create individual with mutated chromosome
		var indiv individual.TestIndividual // var indiv individual.Individual
		output[i] = indiv.Create(child)
	}

	return output // , Statistic.New(population)
}
