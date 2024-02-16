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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Random(tt.args.rng, tt.args.inputSize, tt.args.outputSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Random() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLayer_Propogate(t *testing.T) {
	type args struct {
		inputs []float32
	}
	tests := []struct {
		name        string
		l           *Layer
		args        args
		wantOutputs []float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputs := tt.l.Propogate(tt.args.inputs); !reflect.DeepEqual(gotOutputs, tt.wantOutputs) {
				t.Errorf("Layer.Propogate() = %v, want %v", gotOutputs, tt.wantOutputs)
			}
		})
	}
}
