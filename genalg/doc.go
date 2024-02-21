// Evolve applies the genetic algorithm to the provided population.
// The algorithm consists of the following steps:
//
//  1. Selection: A method is used to select a subset of individuals from the
//     population.
//  2. Crossover: The selected individuals are combined to produce new
//     individuals.
//  3. Mutation: The new individuals are modified to introduce variation.
//  4. Fitness evaluation: The fitness of each individual in the population is
//     re-evaluated.
//  5. Replacement: The new population is replaced by the old population.
//
// The algorithm is repeated for a given number of generations.
//
// The selection method, crossover method, and mutation method are provided
// when creating a new GeneticAlgorithm.
//
// The evaluateFitness function is used to evaluate the fitness of an
// individual. It takes a pointer to an individual and returns a float32
// representing the fitness.
//
// The population is represented as a slice of individuals. The Evolve method
// modifies the contents of this slice.
//
// The number of individuals in the population can be retrieved using the
// Len method.
//
// The individuals in the population can be accessed using the standard
// slice syntax.
//
// The new population is returned.
package genalg
