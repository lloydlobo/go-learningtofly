package layer

import (
	"math/rand"
	"reflect"
	"testing"

	"ffnn/internal/neuron"
)

func TestNew(t *testing.T) {
	type args struct {
		neurons []neuron.Neuron
	}
	tests := []struct {
		name string
		args args
		want Layer
	}{
		{
			"CreatesLayerWithNeurons",
			args{
				[]neuron.Neuron{
					neuron.Random(rand.New(rand.NewSource(0)), uint(4)),
					neuron.Random(rand.New(rand.NewSource(1)), uint(8)),
					neuron.Random(rand.New(rand.NewSource(2)), uint(9)),
				},
			},
			Layer{
				[]neuron.Neuron{
					neuron.Random(rand.New(rand.NewSource(0)), uint(4)),
					neuron.Random(rand.New(rand.NewSource(1)), uint(8)),
					neuron.Random(rand.New(rand.NewSource(2)), uint(9)),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.neurons); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	type args struct {
		rng        *rand.Rand
		inputSize  uint
		outputSize uint
	}
	tests := []struct {
		name string
		args args
		want Layer
	}{
		{
			"CreatesRandomLayer",
			args{rand.New(rand.NewSource(0)), uint(4), uint(3)},
			Layer{[]neuron.Neuron{
				{Bias: 0.8903923, Weights: []float32{-0.51006985, 0.31191254, -0.8913123, -0.26482558}},
				{Bias: -0.4210391, Weights: []float32{-0.6151228, 0.3106643, 0.7943394, -0.6652911}},
				{Bias: -0.42282867, Weights: []float32{0.80520964, 0.6995605, -0.45390642, 0.21816039}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Random(tt.args.rng, tt.args.inputSize, tt.args.outputSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Random() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLayer_Propagate(t *testing.T) {
	n1 := (neuron.New(0.0, []float32{0.1, 0.2, 0.3}))
	n2 := (neuron.New(0.0, []float32{0.4, 0.5, 0.6}))
	lyr := &Layer{[]neuron.Neuron{n1, n2}}

	inputs := []float32{-0.5, 0.0, 0.5} // Note: how did the author infer it?

	type args struct {
		inputs []float32
	}
	tests := []struct {
		name        string
		l           *Layer
		args        args
		wantOutputs []float32
	}{
		{
			"PropagatesInputsThroughNeurons",
			lyr,
			args{inputs},
			[]float32{n1.Propagate(&inputs), n2.Propagate(&inputs)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputs := tt.l.Propagate(tt.args.inputs); !reflect.DeepEqual(gotOutputs, tt.wantOutputs) {
				t.Errorf("Layer.Propagate() = %v, want %v", gotOutputs, tt.wantOutputs)
			}
		})
	}
}
