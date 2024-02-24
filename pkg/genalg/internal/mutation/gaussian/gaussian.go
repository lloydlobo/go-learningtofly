// Package gaussian implements gaussian mutation for package mutation.
package gaussian

import (
	"fmt"
	"math/rand"

	"genalg/internal/chromosome"
)

type GaussianMutation struct {
	Chance float32
	Coeff  float32
}

func New(chance float32, coeff float32) *GaussianMutation {
	if inRange := chance >= 0.0 && chance <= 1.0; !inRange { // err if not in range
		panic(fmt.Errorf("invalid chance: got: %v, want: (0.0..=1.0)", chance))
	}

	return &GaussianMutation{Chance: chance, Coeff: coeff}
}

func (gm GaussianMutation) Mutate(rng *rand.Rand, child *chromosome.Chromosome) {

	for gene := range child.Genes {
		var sign float32
		if rng.Float32() < 0.5 {
			sign = -1.0
		} else {
			sign = 1.0
		}

		// Note: sample = NormFloat64() * desiredStdDev + desiredMean
		//     child.Genes[gene] = child.Genes[gene] + rng.NormFloat32()*gm.Coeff
		if rng.Float32() < gm.Chance {
			child.Genes[gene] += sign * gm.Coeff * rng.Float32()
		}
	}
}
