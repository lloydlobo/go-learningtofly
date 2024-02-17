package ffnn

import (
	"math/rand"
	"reflect"
	"testing"

	"ffnn/internal/layer"
	"ffnn/internal/layertopology"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.layers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetwork_Random(t *testing.T) {
	type args struct {
		rng    *rand.Rand
		layers []layertopology.LayerTopology
	}
	tests := []struct {
		name string
		n    *Network
		args args
		want Network
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Random(tt.args.rng, tt.args.layers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Network.Random() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetwork_Propogate(t *testing.T) {
	type args struct {
		inputs []float32
	}
	tests := []struct {
		name string
		n    *Network
		args args
		want []float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Propagate(tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Network.Propogate() = %v, want %v", got, tt.want)
			}
		})
	}
}
