package neuron

import (
	"math"
	"math/rand"
	"testing"

	"ffnn/internal/testutil"
)

func TestRandom(t *testing.T) {
	var (
		minBias   = float32(-1.0)
		maxBias   = float32(1.0)
		tolerance = testutil.Float32Epsilon // use tolerance as random float32 is involved
	)

	type args struct {
		rng       *rand.Rand
		inputSize uint
	}
	tests := []struct {
		name string
		args args
		want Neuron
	}{
		{
			"ApproxRelativeEq seed src int64(0)",
			args{rand.New(rand.NewSource(0)), 4},
			Neuron{0.8903923, []float32{-0.51006985, 0.31191254, -0.8913123, -0.26482558}},
		},
		{
			"ApproxRelativeEq seed src int64(1)",
			args{rand.New(rand.NewSource(1)), 4},
			Neuron{0.20932055, []float32{0.88101816, 0.32912016, -0.12457162, -0.150725}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Random(tt.args.rng, tt.args.inputSize)

			if got.Bias < minBias || got.Bias > maxBias {
				t.Errorf("generated bias out of range: got: %v, want: [%v, %v]", got.Bias, minBias, maxBias)
			}

			diff := math.Abs(float64(got.Bias - tt.want.Bias))

			if !testutil.ApproxRelativeEq(got.Bias, tt.want.Bias, tolerance) {
				t.Errorf("Random().Bias = %v, want %v, diff: %v, tolerance: %v", got.Bias, tt.want.Bias, diff, tolerance)
			}

			for i, gw := range got.Weights {
				ww := tt.want.Weights[i]
				diff := math.Abs(float64(gw - ww))

				if !testutil.ApproxRelativeEq(gw, ww, tolerance) {
					t.Errorf("Random().Weights[%d] = %v, want %v, diff: %v, tolerance: %v", i, gw, ww, diff, tolerance)
				}
			}
		})
	}
}

func TestNeuron_Propogate(t *testing.T) {
	nron := Neuron{Bias: 0.5, Weights: []float32{-0.3, 0.8}}

	type args struct {
		inputs *[]float32 // ensure equal quantity to avoid triggering panic in Neuron.Propogate function
	}
	tests := []struct {
		name string
		n    *Neuron
		args args
		want float32
	}{
		{
			"`0.5` and `1.0` chosen by a fair dice roll",
			&nron,
			args{&[]float32{0.5, 1.0}},
			float32((-0.3 * 0.5) + (0.8 * 1.0) + 0.5), // Could've written `1.15` right away, but showing the entire formula makes the intentions clearer.
		},
		{
			"max() (ReLU) works",
			&nron,
			args{&[]float32{-10.0, -10.0}},
			float32(0.0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Propagate(tt.args.inputs); got != tt.want {
				t.Errorf("Neuron.Propogate() = %v, want %v", got, tt.want)
			}
		})
	}

	// n := New(0.0, []float32{.1, .2, .3})
	// n.Propogate()
}
