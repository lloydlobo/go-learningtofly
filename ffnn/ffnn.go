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
type Network struct {
	Layers []layer.Layer
}

// New creates a new forward-feed neural network.
func New(layers []layer.Layer) Network {
	return Network{
		Layers: layers,
	}
}

func (n *Network) Random(rng *rand.Rand, layers []layertopology.LayerTopology) Network {
	var builtLayers []layer.Layer // PERF: make with capacity

	for i := range len(layers) - 1 {
		inputSize, outputSize := layers[i].Neurons, layers[i+1].Neurons

		builtLayers = append(builtLayers, layer.Random(rng, inputSize, outputSize))
	}

	return Network{
		Layers: builtLayers,
	}
}

func (n *Network) Propogate(inputs []float32) []float32 {
	for _, l := range n.Layers {
		inputs = l.Propogate(inputs)
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
