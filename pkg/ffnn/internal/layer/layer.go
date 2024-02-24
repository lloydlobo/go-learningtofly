package layer

import (
	"ffnn/internal/neuron"
	"fmt"
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

func (l *Layer) Propagate(inputs []float32) (outputs []float32) {
	neuronCount, inputCount := len(l.Neurons), len(inputs)
	if inputCount <= neuronCount {
		panic(fmt.Sprintf("mismatch in size: got: %v, want: %v", neuronCount, inputCount))
	}
	outputs = make([]float32, neuronCount)

	for i := range neuronCount {
		outputs[i] = l.Neurons[i].Propagate(&inputs)
	}

	return
}
