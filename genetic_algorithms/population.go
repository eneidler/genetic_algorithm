package genetic_algorithm

type Population[T Individual] struct {
	individuals []T
	size        int
}

// NewPopulation creates a new population with a size based on the 'size' parameter.
// Provide a factory method for creating random individuals for the createIndividual func.
func NewPopulation[T Individual](size int, createIndividual func() T) *Population[T] {
	population := &Population[T]{
		individuals: make([]T, size),
		size:        size,
	}

	for i := 0; i < size; i++ {
		population.individuals[i] = createIndividual()
	}

	return population
}

func (p *Population[T]) GetIndividuals() []T {
	return p.individuals
}

func (p *Population[T]) ReplaceIndividuals(individuals []T) {
	p.individuals = individuals
	p.size = len(individuals)
}
