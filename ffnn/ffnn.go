package ffnn

import (
	"fmt"
	"math/rand"

	"ffnn/internal/layer"
	"ffnn/internal/layertopology"
	"ffnn/internal/neuron"
)

// Inside a FFNN, all layers are connected consecutively, back to back.
//
// Network is built from layers:
//
//	o          |          |          |
//	           |  o       |          |
//	o          |          | o        |
//	           |  o       |          |
//	o          |          |          |
//
//	layers[0]  layers[1]  layers[2]
//	   3      ->   2     ->     1
//
// A neural network's most crucial operation is propogating numbers as such:
//
//	f([-0.3, 0.7, 0.0]) = -1.0
//
// .
type Network struct {
	// Note: Layer and Neuron will remain an implementation detail, so we can
	// introduce changes to our implementations without imposing breaking changes
	// on the downstream packages (i.e. our library's users.)

	// .
	layers []layer.Layer
}

// New creates a new forward-feed neural network.
func New(layers []layer.Layer) Network {
	return Network{layers}
}

func Random(rng *rand.Rand, layers []layertopology.LayerTopology) Network {
	builtLayers := make([]layer.Layer, len(layers))

	for i := range len(layers) - 1 {
		inputSize, outputSize := layers[i].Neurons, layers[i+1].Neurons
		builtLayers[i] = layer.Random(rng, inputSize, outputSize)
	}

	return Network{
		layers: builtLayers,
	}
}

func (n *Network) Propagate(inputs []float32) []float32 {
	for _, l := range n.layers {
		inputs = l.Propagate(inputs)
	}

	return inputs
}

func Example() {
	var bias float32 = 0.2
	weights := []float32{0.1, 0.2}

	nron := neuron.New(bias, weights)
	nrons := []neuron.Neuron{nron}

	lyr := layer.New(nrons)
	layers := []layer.Layer{lyr}

	ntwrk := New(layers)

	fmt.Printf("ntwrk: %v\n", ntwrk)
}
