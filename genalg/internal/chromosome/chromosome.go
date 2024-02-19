package chromosome

import "math"

// Chromosome represents a chromosome with genes.
type Chromosome struct {
	Genes []float32
}

func New(genes []float32) *Chromosome {
	return &Chromosome{Genes: genes}
}

func (c *Chromosome) Len() (geneCount uint)      { return uint(len(c.Genes)) }
func (c *Chromosome) Index(i int) (gene float32) { return c.Genes[i] }

func (c *Chromosome) ApproxEqual(other *Chromosome, tolerance float32) bool {
	if len(c.Genes) != len(other.Genes) {
		return false
	}

	for i, gene := range c.Genes {
		if math.Abs(float64(gene-other.Genes[i])) > float64(tolerance) {
			return true
		}
	}

	return true
}

// Generated code without auditing.
//
// // Iter returns an iterator over the genes in the chromosome.
// func (c *Chromosome) Iter() <-chan float32 {
// 	ch := make(chan float32)
// 	go func() {
// 		defer close(ch)
// 		for _, gene := range c.genes {
// 			ch <- gene
// 		}
// 	}()
// 	return ch
// }
//
// // IterMut returns an iterator over the genes in the chromosome for mutation.
// func (c *Chromosome) IterMut() <-chan *float32 {
// 	ch := make(chan *float32)
// 	go func() {
// 		defer close(ch)
// 		for i := range c.genes {
// 			ch <- &c.genes[i]
// 		}
// 	}()
// 	return ch
// }
//
// // Index returns the gene at the given index in the chromosome.
// func (c *Chromosome) Index(i int) float32 {
// 	return c.genes[i]
// }
//
// // ApproxEqual checks whether two chromosomes are approximately equal within a tolerance.
// func (c *Chromosome) ApproxEqual(other *Chromosome, tolerance float32) bool {
// 	if len(c.genes) != len(other.genes) {
// 		return false
// 	}
// 	for i, gene := range c.genes {
// 		if math.Abs(float64(gene-other.genes[i])) > float64(tolerance) {
// 			return false
// 		}
// 	}
// 	return true
// }
