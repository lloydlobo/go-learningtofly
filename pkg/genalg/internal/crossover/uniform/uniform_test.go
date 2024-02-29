package uniform

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"genalg/internal/chromosome"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *UniformCrossover
	}{
		{"Create New UniformCrossover", &UniformCrossover{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
FIXME: Errorlog

	=== RUN   TestUniformCrossover_Crossover_Diff
	--- FAIL: TestUniformCrossover_Crossover_Diff (0.00s)
	panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	        panic: runtime error: invalid memory address or nil pointer dereference
	[signal 0xc0000005 code=0x0 addr=0x0 pc=0xb22617]
*/
func TestUniformCrossover_Crossover_Diff(t *testing.T) {
	rng := rand.New(rand.NewSource(0))

	var (
		parentA chromosome.Chromosome
		parentB chromosome.Chromosome
	)

	const geneCount = 100
	for i := 0; i < geneCount; i++ {
		parentA.Genes = append(parentA.Genes, float32(i+1))
		parentB.Genes = append(parentB.Genes, -float32(i+1))
	}
	fmt.Printf("parentA: %v\n", parentA)
	fmt.Printf("parentB: %v\n", parentB)

	uc := UniformCrossover{}
	child := uc.Crossover(rng, parentA, parentB)
	fmt.Printf("child: %v\n", child)

	// Number of genes different between `child` and `parentA`
	var diffA int
	for i, c := range child.Genes {
		if c != parentA.Index(i) {
			diffA++
		}
	}

	// Number of genes different between `child` and `parentB`
	var diffB int
	for i, c := range child.Genes {
		if c != parentB.Index(i) {
			diffB++
		}
	}

	if diffA != 47 {
		t.Errorf("Number of genes different between `child` and `parentA`: %v", diffA)
	}
	if diffB != 53 {
		t.Errorf("Number of genes different between `child` and `parentB`: %v", diffB)
	}
}
