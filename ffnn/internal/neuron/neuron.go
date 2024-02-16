package neuron

import (
	"math/rand"
)

type Neuron struct {
	Bias    float32
	Weights []float32
}

// # Errors
//
//   - Panics if weights are empty.
func New(bias float32, weights []float32) Neuron {
	if len(weights) == 0 {
		panic("weights must not be empty")
	}
	return Neuron{Bias: bias, Weights: weights}
}

// Random creates a new neuron with random bias and weights.
func Random(rng *rand.Rand, inputSize uint) Neuron {

	// Generate a float32 pseudo-random number in the half-open interval [0.0,1.0)
	// and then scale and shift to fit the range -1.0 to 1.0.
	bias := (rng.Float32() * 2) - 1

	weights := make([]float32, inputSize)

	for i := range weights {
		weights[i] = (rng.Float32() * 2) - 1
	}

	return New(bias, weights)
}

// # Errors
//
//   - Panics if count of inputs and Neuron.Weights do not match.
func (n *Neuron) Propogate(inputs *[]float32) float32 {
	count := len(*inputs)

	if count != len(n.Weights) {
		panic("len(inputs)!=len(n.Weights)")
	}

	var output float32

	for i := range count {
		output += (*inputs)[i] * n.Weights[i]
	}

	return max(0.0, n.Bias+output)
}
