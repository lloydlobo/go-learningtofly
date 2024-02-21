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
				ChromosomeValue: chromosome.Chromosome{},
				FitnessValue:    0.5,
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
				ChromosomeValue: chromosome.Chromosome{Genes: []float32{0.1, 0.2, 0.3}},
				FitnessValue:    0,
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
			ti:   TestIndividual{ChromosomeValue: chromosome.Chromosome{Genes: []float32{0.1, 0.2, 0.3}}},
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
			ti:   TestIndividual{FitnessValue: 0.5},
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

func TestNewTestIndividualWithChromosome(t *testing.T) {
	type args struct {
		chromosome chromosome.Chromosome
	}
	tests := []struct {
		name string
		args args
		want Individual
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTestIndividualWithChromosome(tt.args.chromosome); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTestIndividualWithChromosome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestIndividualWithChromosome_Create(t *testing.T) {
	type args struct {
		chromosome chromosome.Chromosome
	}
	tests := []struct {
		name string
		tr   TestIndividualWithChromosome
		args args
		want Individual
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Create(tt.args.chromosome); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestIndividualWithChromosome.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestIndividualWithChromosome_Chromosome(t *testing.T) {
	tests := []struct {
		name string
		iwc  TestIndividualWithChromosome
		want *chromosome.Chromosome
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.iwc.Chromosome(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestIndividualWithChromosome.Chromosome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestIndividualWithChromosome_Fitness(t *testing.T) {
	tests := []struct {
		name string
		iwc  TestIndividualWithChromosome
		want float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.iwc.Fitness(); got != tt.want {
				t.Errorf("TestIndividualWithChromosome.Fitness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTestIndividualWithFitness(t *testing.T) {
	type args struct {
		fitness float32
	}
	tests := []struct {
		name string
		args args
		want Individual
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTestIndividualWithFitness(tt.args.fitness); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTestIndividualWithFitness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestIndividualWithFitness_Create(t *testing.T) {
	type args struct {
		chromosome chromosome.Chromosome
	}
	tests := []struct {
		name string
		tr   TestIndividualWithFitness
		args args
		want Individual
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Create(tt.args.chromosome); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestIndividualWithFitness.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestIndividualWithFitness_Chromosome(t *testing.T) {
	tests := []struct {
		name string
		iwf  TestIndividualWithFitness
		want *chromosome.Chromosome
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.iwf.Chromosome(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestIndividualWithFitness.Chromosome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestIndividualWithFitness_Fitness(t *testing.T) {
	tests := []struct {
		name string
		iwf  TestIndividualWithFitness
		want float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.iwf.Fitness(); got != tt.want {
				t.Errorf("TestIndividualWithFitness.Fitness() = %v, want %v", got, tt.want)
			}
		})
	}
}
