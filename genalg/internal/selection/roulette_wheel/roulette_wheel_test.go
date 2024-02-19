package roulettewheel

import (
	"math/rand"
	"reflect"
	"testing"

	"genalg/internal/individual"
)

func TestRouletteWheelSelection_Select(t *testing.T) {
	type args struct {
		rng        *rand.Rand
		population *[]individual.Individual
	}
	tests := []struct {
		name string
		r    *RouletteWheelSelection
		args args
		want *individual.Individual
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Select(tt.args.rng, tt.args.population); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouletteWheelSelection.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
