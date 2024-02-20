// Package uniform implements methods for crossover.
package uniform

import (
	"fmt"
	"math/rand"

	"genalg/internal/chromosome"
)

// UniformCrossover represents the uniform crossover method.
type UniformCrossover struct{}

func New() *UniformCrossover {
	return &UniformCrossover{}
}

// Crossover performs uniform crossover between two parent chromosomes.
func (uc *UniformCrossover) Crossover(
	rng *rand.Rand,
	parentA,
	parentB *chromosome.Chromosome,
) chromosome.Chromosome {
	if na, nb := parentA.Len(), parentB.Len(); na != nb {
		panic(fmt.Sprintf("parent chromosomes must have the same length: a: %v, b: %v", na, nb))
	}

	child := make([]float32, parentA.Len()) // child genes
	const probability = 0.5                 // 50/50 chance

	for i, geneA := range parentA.Genes { // Note: geneA value for perf via range
		if rng.Float32() < probability {
			child[i] = geneA
		} else {
			child[i] = parentB.Index(i)
		}
	}

	return chromosome.New(child)
}
