package ga

type Individual interface {
	CalculateFitness() float64

	Crossover(parentB Individual) []Individual

	Mutate()

	Fitness() float64

	Clone() Individual
}
