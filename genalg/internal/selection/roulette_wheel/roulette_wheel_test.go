package roulettewheel

import (
	"math/rand"
	"reflect"
	"testing"

	"genalg/internal/individual"
)

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
			var pop []individual.Individual // reflect.ValueOf(population)
			for _, indiv := range *tt.args.population {
				pop = append(pop, indiv)
			}
			if got := tt.r.Select(tt.args.rng, &pop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouletteWheelSelection.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

	/*
		Notice that while building the histogram, we're casting fitness scores from f32 to i32:
		We have to do that, because floating-point numbers in Rust don't implement the Ord trait,
		making it impossible to use them as a BTreeMap-'s key:
		The reason is that floating-point numbers, as defined by the IEEE 754 standard,
		are not a totally ordered set - namely, comparing NaN-s is problematic, because:
		NaN != NaN
		Practically, it means that had you ever inserted a NaN into a map,
		not only would you be unable to retrieve that particular entry
		using .get(), but you could break BTreeMap-'s internal structure,
		making it impossible to retrieve any item.
		(by the way, that's also true for custom implementations of
		Ord and PartialOrd - if they don't satisfy asymmetry and transitivity,
		you're gonna have a bad time.)
	*/
	for range iterations {
		selected := method.Select(rng, &popul)
		actualHistogram[int32(selected.Fitness())]++
	}

	expectedHistogram := map[int32]int{
		// (fitness, how many times this fitness has been chosen)
		1: 98,
		2: 202,
		3: 297, // 278
		4: 403, // 422
	} // Note: values may differ due to rand src in rng

	for key, val := range expectedHistogram {
		if got := actualHistogram[int32(key)]; got != val {
			t.Errorf("histograms do not match at key %v: got: %v, want: %v", key, got, val)
			break
		}
	}
}
