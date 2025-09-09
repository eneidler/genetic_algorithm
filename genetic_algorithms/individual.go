package genetic_algorithm

type Individual interface {
	CalculateFitness() float64

	Crossover(parentB Individual) []Individual

	Mutate(mutationRate float64)

	Fitness() float64

	Clone() Individual
}
