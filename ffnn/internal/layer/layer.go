package layer

import (
	"ffnn/internal/neuron"
	"math/rand"
)

type Layer struct {
	Neurons []neuron.Neuron
}

func New(neurons []neuron.Neuron) Layer {
	return Layer{Neurons: neurons}
}

func Random(rng *rand.Rand, inputSize, outputSize uint) Layer {
	neurons := make([]neuron.Neuron, outputSize)

	for i := range outputSize {
		neurons[i] = neuron.Random(rng, inputSize)
	}

	return Layer{Neurons: neurons}
}

func (l *Layer) Propogate(inputs []float32) (outputs []float32) {
	outputs = make([]float32, len(inputs))

	for i, n := range l.Neurons {
		outputs[i] = n.Propogate(&inputs)
	}

	return
}
