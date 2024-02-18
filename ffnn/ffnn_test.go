package ffnn

import (
	"math/rand"
	"reflect"
	"testing"

	"ffnn/internal/layer"
	"ffnn/internal/layertopology"
	"ffnn/internal/neuron"
)

func TestExample(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Example function panicked: %v", r)
		}
	}()

	Example()
}

func TestNew(t *testing.T) {
	type args struct {
		layers []layer.Layer
	}
	tests := []struct {
		name string
		args args
		want Network
	}{
		{
			"CreateNetworkWithOneLayer",
			args{layers: []layer.Layer{layer.New([]neuron.Neuron{neuron.New(0.1, []float32{0.2, 0.3})})}},
			Network{layers: []layer.Layer{{Neurons: []neuron.Neuron{{Bias: 0.1, Weights: []float32{0.2, 0.3}}}}}},
		},
		{
			"CreateNetworkWithTwoLayers",
			args{layers: []layer.Layer{
				layer.New([]neuron.Neuron{neuron.New(0.1, []float32{0.2, 0.3})}),
				layer.New([]neuron.Neuron{neuron.New(0.2, []float32{0.4, 0.5})}),
			}},
			Network{layers: []layer.Layer{
				{Neurons: []neuron.Neuron{{Bias: 0.1, Weights: []float32{0.2, 0.3}}}},
				{Neurons: []neuron.Neuron{{Bias: 0.2, Weights: []float32{0.4, 0.5}}}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.layers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(0))

	type args struct {
		rng    *rand.Rand
		layers []layertopology.LayerTopology
	}
	tests := []struct {
		name string
		args args
		want Network
	}{
		{
			// TODO: when len(layers) <= 1, expect error.
			//	This feature is currently disabled via flag.
			"CreateRandom_EmptyNetwork_LayerTopology",
			args{rng, []layertopology.LayerTopology{}},
			Network{
				layers: []layer.Layer{},
			},
		},
		{
			"CreateRandom_Network_LayerTopology",
			args{rng, []layertopology.LayerTopology{{Neurons: 3}, {Neurons: 2}, {Neurons: 1}}},
			Network{
				layers: []layer.Layer{
					{Neurons: []neuron.Neuron{
						{Bias: 0.8903923, Weights: []float32{-0.51006985, 0.31191254, -0.8913123}},
						{Bias: -0.26482558, Weights: []float32{-0.4210391, -0.6151228, 0.3106643}}}}, // layer with 3 neurons
					{Neurons: []neuron.Neuron{
						{Bias: 0.7943394, Weights: []float32{-0.6652911, -0.42282867}}}}, // layer with 2 neurons
					{}, // Empty layer
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Random(tt.args.rng, tt.args.layers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Random() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetwork_Propogate(t *testing.T) {
	rng := rand.New(rand.NewSource(0))

	emptyLayers := []layertopology.LayerTopology{}
	emptyNetwork := Random(rng, emptyLayers)

	layers := []layer.Layer{
		{Neurons: []neuron.Neuron{neuron.New(0.0, []float32{-0.5, -0.4, -0.3}), neuron.New(0.0, []float32{-0.2, -0.1, 0.0})}},
		{Neurons: []neuron.Neuron{neuron.New(0.0, []float32{-0.5, -0.5})}},
	} // or use `layers := []layertopology.LayerTopology{{Neurons: 3}, {Neurons: 2}, {Neurons: 1}}`
	network := New(layers)

	type args struct {
		inputs []float32
	}
	tests := []struct {
		name string
		n    *Network
		args args
		want []float32
	}{
		{
			"PropagateInputsThroughEmptyLayers",
			&emptyNetwork,
			args{[]float32{}},
			[]float32{},
		},
		{ // See reference: https://github.com/Patryk27/shorelark/blob/d598ef91f250db870af285c0f433d976170d649f/libs/neural-network/src/lib.rs#L144
			"PropagateInputsThroughLayers",
			&network,
			args{[]float32{0.5, 0.6, 0.7}},
			layers[1].Propagate(layers[0].Propagate([]float32{0.5, 0.6, 0.7})),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Propagate(tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Network.Propogate() = %v, want %v", got, tt.want)
			}
		})
	}
}
