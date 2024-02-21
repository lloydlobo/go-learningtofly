package genalg

import (
	"genalg/internal/chromosome"
	"genalg/internal/crossover"
	"genalg/internal/crossover/uniform"
	"genalg/internal/individual"
	"genalg/internal/mutation"
	"genalg/internal/mutation/gaussian"
	roulettewheel "genalg/internal/selection/roulette_wheel"
	"math/rand"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type S SelectionMethod

	type args struct {
		selector S
		crosser  crossover.CrossoverMethod
		mutator  mutation.MutationMethod
	}
	tests := []struct {
		name string
		args args
		want GeneticAlgorithm
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.selector, tt.args.crosser, tt.args.mutator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
	 Errorlogs:

		- panic: test timed out after 30s
			>>> Could it be the while loop? (NO)
			>>> Only this is the failing test:
					=== RUN   TestUniformCrossover_Crossover_Diff
					--- FAIL: TestUniformCrossover_Crossover_Diff (0.00s)
					panic: runtime error: invalid memory address or nil pointer dereference [recovered]
							panic: runtime error: invalid memory address or nil pointer dereference
					[signal 0xc0000005 code=0x0 addr=0x0 pc=0xb22617]
			>> TODO: Make sure we check for more than 1 parent layers
		-	--- FAIL: TestGeneticAlgorithm_Evolve (0.00s)
			panic: runtime error: invalid memory address or nil pointer dereference [recovered]
			        panic: runtime error: invalid memory address or nil pointer dereference
			[signal 0xc0000005 code=0x0 addr=0x18 pc=0x3d2470]
*/
func TestGeneticAlgorithm_Evolve(t *testing.T) {
	// Helper closure to create a few individuals at once.
	newIndividual := func(genes []float32) (indiv individual.TestIndividual) {
		var chromo chromosome.Chromosome
		chromo.Genes = append(chromo.Genes, genes...)
		return indiv.Create(chromo).(individual.TestIndividual) // coerce type
	}

	const epocs int = 10 // try atleast 10 epocs/generations

	selector := roulettewheel.RouletteWheelSelection{}
	crosser := uniform.UniformCrossover{}
	mutator := gaussian.GaussianMutation{Chance: 0.5, Coeff: 0.5}

	ga := New(selector, crosser, mutator)

	rng := rand.New(rand.NewSource(0))

	type args struct {
		rng        *rand.Rand
		population []individual.Individual
		fitnessFn  func(indiv individual.Individual) float32
	}
	tests := []struct {
		name string
		ga   GeneticAlgorithm
		args args
		want []individual.Individual
	}{
		{
			name: "",
			ga:   ga,
			args: args{
				rng: rng,
				population: []individual.Individual{
					// FIXME: author's fitness for these individual may be achieved via TestIndividual enums setup properly.
					// 		it's 0.0 for all right now.
					newIndividual([]float32{0.0, 0.0, 0.0}), // fitness = 0.0
					newIndividual([]float32{1.0, 1.0, 1.0}), // fitness = 3.0
					newIndividual([]float32{1.0, 2.0, 1.0}), // fitness = 4.0
					newIndividual([]float32{1.0, 2.0, 4.0}), // fitness = 7.0
				},
				fitnessFn: func(indiv individual.Individual) float32 { return 0.0 }, // stub unimplemented
			},
			// Note: using zeroed-out expected population initially. Also individual.Create does not take fitness as arg so expect zero-value fitness.
			// NOTE: solved by using GaussianMutation{Chance:0.5,Coeff:0.5}
			want: []individual.Individual{
				newIndividual([]float32{-0.30737132, 0.038694657, -0.24632761}),
				newIndividual([]float32{-0.30737132, 0.32855722, 0.11804612}),
				newIndividual([]float32{-0.69411385, 0.22006679, 0.60720074}),
				newIndividual([]float32{-0.30737132, 0.34553945, 0.25269297}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for range epocs {
				if dbgEnabled := false; dbgEnabled {
					for _, popul := range tt.args.population {
						fitness := popul.Fitness()
						t.Logf("fitness: %v\n", fitness)
					}
				}
				next := tt.ga.Evolve(tt.args.rng, tt.args.population, tt.args.fitnessFn)
				tt.args.population = next
			}
			if got := tt.args.population; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneticAlgorithm.Evolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
