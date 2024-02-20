package individual

import (
	"genalg/internal/chromosome"
)

type Individual interface {
	Create(chromosome chromosome.Chromosome) Individual
	Chromosome() *chromosome.Chromosome
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
	chromosome chromosome.Chromosome // Note: add this later
	fitness    float32
}

// NewTestIndividual creates a new TestIndividual with the given fitness
func NewTestIndividual(fitness float32) *TestIndividual {
	return &TestIndividual{chromosome: chromosome.Chromosome{}, fitness: fitness} // chromosome is not set in initializer?
}

func (TestIndividual) Create(chromosome chromosome.Chromosome) Individual {
	return TestIndividual{chromosome: chromosome, fitness: 0} // fitness not set in Create method
}

func (ti TestIndividual) Chromosome() *chromosome.Chromosome { return &ti.chromosome }
func (ti TestIndividual) Fitness() float32                   { return ti.fitness }
