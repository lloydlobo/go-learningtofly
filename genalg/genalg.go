package genalg

import (
	"fmt"
)

type GeneticAlgorithm[S any] struct {
	selectionMethod S

	// crossoverMethod CrossoverMethod[I]

	// mutationMethod MutationMethod[I]
}

func New[S any](selectionMethod S) GeneticAlgorithm[S] {
	return GeneticAlgorithm[S]{
		selectionMethod: selectionMethod,
	}
}

func (ga GeneticAlgorithm[I]) Evolve(
	population *[]I,
	evaluateFitness func(individual *I) float32,
) []I {

	populCount := len(*population)

	if populCount == 0 {
		panic("expected population to not be empty")
	}
	fmt.Printf("ga.selectionMethod: %v\n", ga.selectionMethod)

	output := make([]I, populCount)

	for i, p := range *population {
		for j := i + 1; j < populCount; j++ {
			fmt.Printf("[]interface { i j p }: {%v, %v, %v}\n", i, j, p)
			// TODO selection
			// TODO crossover
			// TODO mutation
			output[i] = p
		}
	}

	return output
}
