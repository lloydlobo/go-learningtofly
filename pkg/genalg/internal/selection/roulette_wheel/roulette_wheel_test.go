package roulettewheel

import (
	"math/rand"
	"reflect"
	"testing"

	"genalg/internal/individual"
)

func TestRouletteWheelSelection_Select_WithHistogram(t *testing.T) {
	method := RouletteWheelSelection{}
	rng := rand.New(rand.NewSource(0))

	population := []individual.TestIndividual{
		*individual.NewTestIndividual(2.0),
		*individual.NewTestIndividual(1.0),
		*individual.NewTestIndividual(4.0),
		*individual.NewTestIndividual(3.0),
	}

	const iterations = 1000
	actualHistogram := make(map[int32]int)

	// Note: To avoid panic: interface conversion: individual.Individual is
	// individual.TestIndividual, not *individual.TestIndividual [recovered]
	// Coercing type to satisfy without using generics in Select method
	var popul []individual.Individual
	for _, p := range population {
		popul = append(popul, p)
	}

	for i := 0; i < iterations; i++ {
		selected := method.Select(rng, &popul)
		actualHistogram[int32(selected.Fitness())]++
	}

	expectedHistogram := map[int32]int{
		// (fitness, how many times this fitness has been chosen)
		1: 113, // 98,
		2: 173, // 202,
		3: 302, // 297, // 278
		4: 412, // 403, // 422
	} // Note: values may differ due to rand src in rng

	for key, val := range expectedHistogram {
		if got := actualHistogram[int32(key)]; got != val {
			t.Errorf("histograms do not match at key %v: got: %v, want: %v", key, got, val)
			break
		}
	}
}

func TestRouletteWheelSelection_Select(t *testing.T) {
	method := RouletteWheelSelection{}
	rng := rand.New(rand.NewSource(0))

	population := []individual.TestIndividual{
		*individual.NewTestIndividual(2.0),
		*individual.NewTestIndividual(1.0),
		*individual.NewTestIndividual(4.0),
		*individual.NewTestIndividual(3.0),
	}

	type args struct {
		rng        *rand.Rand
		population *[]individual.TestIndividual
	}
	tests := []struct {
		name string
		r    *RouletteWheelSelection
		args args
		want individual.TestIndividual
	}{
		{
			"RouletteWheelSelection_Select method returns a Individual from selection method",
			&method,
			args{rng, &population},
			*individual.NewTestIndividual(4.0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var popul []individual.Individual // reflect.ValueOf(population)
			for _, indiv := range *tt.args.population {
				popul = append(popul, indiv)
			}
			if got := tt.r.Select(tt.args.rng, &popul); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouletteWheelSelection.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
