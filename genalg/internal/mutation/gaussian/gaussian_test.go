package gaussian

import (
	"math/rand"
	"reflect"
	"testing"

	"genalg/internal/chromosome"
)

func actual(chance, coeff float32) (genes []float32) {
	child := chromosome.Chromosome{Genes: []float32{1., 2., 3., 4., 5.}}
	rng := rand.New(rand.NewSource(0))

	New(chance, coeff).Mutate(rng, &child)

	return child.Genes
}

func TestGaussianMutation(t *testing.T) {
	type args struct {
		chance float32
		coeff  float32
	}
	tests := []struct {
		name string
		gm   GaussianMutation
		args args
		want []float32
	}{
		{
			name: "GivenZeroChance_AndZeroCoefficient_DoesNotChangeOriginalChromosome",
			gm:   GaussianMutation{},
			args: args{chance: 0.0, coeff: 0.0},
			want: []float32{1., 2., 3., 4., 5.},
		},
		{
			name: "GivenZeroChance_AndNonZeroCoefficient_DoesNotChangeOriginalChromosome",
			gm:   GaussianMutation{},
			args: args{chance: 0.0, coeff: 0.5},
			want: []float32{1., 2., 3., 4., 5.},
		},
		{
			name: "GivenFiftyFiftyChance_AndZeroCoefficient_DoesNotChangeOriginalChromosome",
			gm:   GaussianMutation{},
			args: args{chance: 0.5, coeff: 0.0},
			want: []float32{1., 2., 3., 4., 5.},
		},
		{
			name: "GivenFiftyFiftyChance_AndNonZeroCoefficient_SlightlyChangesOriginalChromosome",
			gm:   GaussianMutation{},
			args: args{chance: 0.5, coeff: 0.5},
			want: []float32{1.3279781, 1.8552598, 3., 4.144293, 5.},
		},
		{
			name: "GivenMaxChance_AndZeroCoefficient_DoesNotChangeOriginalChromosome",
			gm:   GaussianMutation{},
			args: args{chance: 1.0, coeff: 0.0},
			want: []float32{1., 2., 3., 4., 5.},
		},
		{
			name: "GivenMaxChance_AndNonZeroCoefficient_EntirelyChangesOriginalChromosome",
			gm:   GaussianMutation{},
			args: args{chance: 1.0, coeff: 0.5},
			want: []float32{1.3279781, 1.8552598, 2.5514152, 3.5486975, 5.30454},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := actual(tt.args.chance, tt.args.coeff); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("actual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		chance float32
		coeff  float32
	}
	tests := []struct {
		name string
		args args
		want *GaussianMutation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.chance, tt.args.coeff); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGaussianMutation_Mutate(t *testing.T) {
	type args struct {
		rng   *rand.Rand
		child *chromosome.Chromosome
	}
	tests := []struct {
		name string
		gm   GaussianMutation
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.gm.Mutate(tt.args.rng, tt.args.child)
		})
	}
}
