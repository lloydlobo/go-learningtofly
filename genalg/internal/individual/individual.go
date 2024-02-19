package individual

import (
	"genalg/internal/chromosome"
)

type Individual interface {
	Create(chromosome chromosome.Chromosome) Individual
	Chromosome() *chromosome.Chromosome
	Fitness() float32
}
