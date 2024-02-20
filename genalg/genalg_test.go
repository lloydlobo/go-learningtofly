package genalg

import (
	"reflect"
	"testing"

	"genalg/internal/crossover"
	"genalg/internal/mutation"
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
		want GeneticAlgorithm[S]
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
