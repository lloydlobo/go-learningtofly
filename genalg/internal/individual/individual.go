package individual

import (
	"genalg/internal/chromosome"
)

type Individual interface {
	Create(chromosome chromosome.Chromosome) Individual // individual from chromosome
	Chromosome() *chromosome.Chromosome                 // individual to chromosome
	Fitness() float32
}

// Testing beyond this point...

// TestIndividual is used for testing purposes.
//
// Methods:
//   - Create creates a TestIndividual with the given chromosome.
//   - Chromosome returns the chromosome of the TestIndividual.
//   - Fitness returns the fitness of the TestIndividual.
type TestIndividual struct {
	// (instead of enum, you could also create two separate types like
	// TestIndividualWithChromosome and TestIndividualWithFitness - but that
	// feels too enterprise-y for me, so I'll take a rain check)

	// or use  *chromosome.Chromosome

	chromosome chromosome.Chromosome // Note: add this later
	fitness    float32
}

// NewTestIndividual creates a new TestIndividual with the given fitness
func NewTestIndividual(fitness float32) *TestIndividual {
	return &TestIndividual{
		chromosome: chromosome.Chromosome{}, // FIXME: chromosome is not set in initializer?
		fitness:    fitness,
	}
}

// TODO: using pointer receiver methods for all below creates warning...

// return NewTestIndividual(0)
func (TestIndividual) Create(chromosome chromosome.Chromosome) Individual {
	return TestIndividual{chromosome: chromosome, fitness: 0}
}

func (ti TestIndividual) Chromosome() *chromosome.Chromosome {
	return &ti.chromosome
}

func (ti TestIndividual) Fitness() float32 {
	return ti.fitness
}

// ===========================================================

// Individual enumerations for testing.
type ( // or use a const enum type to match via type of instance?
	TestIndividualWithChromosome struct{ chromosome chromosome.Chromosome }
	TestIndividualWithFitness    struct{ fitness float32 }
)

func NewTestIndividualWithChromosome(chromosome chromosome.Chromosome) Individual {
	return TestIndividualWithChromosome{chromosome}
}
func (TestIndividualWithChromosome) Create(chromosome chromosome.Chromosome) Individual {
	return TestIndividualWithChromosome{chromosome}
}
func (iwc TestIndividualWithChromosome) Chromosome() *chromosome.Chromosome {
	return &iwc.chromosome
}
func (iwc TestIndividualWithChromosome) Fitness() float32 {
	return iwc.chromosome.Sum()
}

func NewTestIndividualWithFitness(fitness float32) Individual {
	return TestIndividualWithFitness{fitness}
}
func (TestIndividualWithFitness) Create(chromosome chromosome.Chromosome) Individual {
	return TestIndividualWithFitness{fitness: 0.0}
}
func (iwf TestIndividualWithFitness) Chromosome() *chromosome.Chromosome {
	panic("not supported for TestIndividualWithFitness")
}
func (iwf TestIndividualWithFitness) Fitness() float32 {
	return iwf.fitness
}
