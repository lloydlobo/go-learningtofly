package ffnn

import (
	"fmt"
	"math/rand"

	"ffnn/internal/layer"
	"ffnn/internal/layertopology"
	"ffnn/internal/neuron"
)

// Note: Layer and Neuron will remain an implementation detail, so we can
// introduce changes to our implementations without imposing breaking changes
// on the downstream packages (i.e. our library's users.)

// Network represents a neural network composed of layers, each containing
// neurons with biases and output weights.
//
// Neurons within layers are connected via synapses or weights, where signals
// are transmitted between neurons. Neurons make decisions on whether to
// propagate signals further or not.
//
// Fields:
//   - layers: A slice of layers composing the network.
//
// Inside a FFNN, all layers are connected consecutively, back to back.
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
//   - f([-0.3, 0.7, 0.0]) = -1.0
type Network struct {
	layers []layer.Layer
}

// New creates a new forward-feed neural network with the given layers.
//
// Parameters:
//   - layers: A slice of layers to compose the network.
//
// Returns:
//   - A new Network instance.
func New(layers []layer.Layer) Network {
	return Network{layers}
}

// Random creates a new forward-feed neural network with
// randomly initialized weights.
//
// Parameters:
//   - rng: A pointer to a rand.Rand instance for generating random numbers.
//   - layers: A slice of layer topologies defining the structure of the network.
//
// Returns:
//   - A new Network instance with randomly initialized weights.
func Random(rng *rand.Rand, layers []layertopology.LayerTopology) Network {
	flagImplemented := false

	layerCount := len(layers)

	if flagImplemented {
		// See reference, https://github.com/Patryk27/shorelark/blob/d598ef91f250db870af285c0f433d976170d649f/libs/neural-network/src/lib.rs#L24
		if layerCount <= 1 {
			panic(fmt.Sprintf("expected layers [%+v] to have more than one layer, got: %v", layers, layerCount))
		}
	}

	builtLayers := make([]layer.Layer, layerCount)

	for i := range len(layers) - 1 {
		inputSize, outputSize := layers[i].Neurons, layers[i+1].Neurons
		builtLayers[i] = layer.Random(rng, inputSize, outputSize)
	}

	return Network{
		layers: builtLayers,
	}
}

// Propagate propagates the input through the network and returns the output.
//
// Parameters:
//   - inputs: A slice of float32 representing the input data to
//     propagate through the network.
//
// Returns:
//   - A slice of float32 representing the output data propagated
//     through the network.
//
// Analogy:
//
//   - Neurons within layers can be compared to biological neurons connected
//     via synapses.
//   - Signals are carried between neurons, with each neuron making decisions
//     based on the received signals to propagate them further or
//     stop their propagation.
//
// Reference: https://pwy.io/posts/learning-to-fly-pt2/#coding-propagate
func (n *Network) Propagate(inputs []float32) []float32 {
	for _, l := range n.layers {
		inputs = l.Propagate(inputs)
	}

	return inputs
}

// Description:
//   - Provides an example demonstrating how to create and use a neural network.
//
// Functionality:
//   - Creates a simple neural network with one neuron and one layer, then prints it.
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
