package individual

import (
	"reflect"
	"testing"

	"genalg/internal/chromosome"
)

func TestNewTestIndividual(t *testing.T) {
	type args struct {
		fitness float32
	}
	tests := []struct {
		name string
		args args
		want *TestIndividual
	}{
		{
			name: "NewTestIndividual creates a TestIndividual with the given fitness",
			args: args{fitness: 0.5},
			want: &TestIndividual{
				chromosome: chromosome.Chromosome{},
				fitness:    0.5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTestIndividual(tt.args.fitness); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTestIndividual() = %v, want %v", got, tt.want)
			}
		})
	}
}

// #[test]
// fn test() {
//     let population = vec![
//         TestIndividual::new(2.0),
//         TestIndividual::new(1.0),
//         TestIndividual::new(4.0),
//         TestIndividual::new(3.0),
//     ];
//
//     /* ... */
// }

func TestTestIndividual_Create(t *testing.T) {
	type args struct {
		chromosome chromosome.Chromosome
	}
	tests := []struct {
		name string
		tr   TestIndividual
		args args
		want Individual
	}{
		{
			name: "TestIndividual Create returns a TestIndividual with the given chromosome",
			tr:   TestIndividual{},
			args: args{chromosome: chromosome.Chromosome{Genes: []float32{0.1, 0.2, 0.3}}},
			want: TestIndividual{
				chromosome: chromosome.Chromosome{Genes: []float32{0.1, 0.2, 0.3}},
				fitness:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Create(tt.args.chromosome); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestIndividual.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestIndividual_Chromosome(t *testing.T) {
	tests := []struct {
		name string
		ti   TestIndividual
		want *chromosome.Chromosome
	}{
		{
			name: "TestIndividual Chromosome returns the chromosome of the TestIndividual",
			ti:   TestIndividual{chromosome: chromosome.Chromosome{Genes: []float32{0.1, 0.2, 0.3}}},
			want: &chromosome.Chromosome{Genes: []float32{0.1, 0.2, 0.3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ti.Chromosome(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestIndividual.Chromosome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestIndividual_Fitness(t *testing.T) {
	tests := []struct {
		name string
		ti   TestIndividual
		want float32
	}{
		{
			name: "TestIndividual Fitness returns the fitness of the TestIndividual",
			ti:   TestIndividual{fitness: 0.5},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ti.Fitness(); got != tt.want {
				t.Errorf("TestIndividual.Fitness() = %v, want %v", got, tt.want)
			}
		})
	}
}
