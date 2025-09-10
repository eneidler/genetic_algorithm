package ga

import (
	"fmt"
)

type Engine[T Individual] struct {
}

func CreateEngine[T Individual](config Config, createIndividual func() T) (*Engine[T], error) {
	// Error handling
	if config.PopulationSize < 2 {
		return nil, fmt.Errorf("population size must be at least 2. provided size: %d", config.PopulationSize)
	}
	if config.MutationRate < 0 || config.MutationRate > 1 {
		return nil, fmt.Errorf("mutation rate must be between 0 and 1, provided rate: %f", config.MutationRate)
	}
	if config.CrossoverRate < 0 || config.CrossoverRate > 1 {
		return nil, fmt.Errorf("crossover rate must be between 0 and 1, provided rate: %f", config.CrossoverRate)
	}
	if config.Generations < 1 {
		return nil, fmt.Errorf("generations must be at least 1. provided generations: %d", config.Generations)
	}

}

func (e *Engine[T]) Run() (T, float64, error) {

}
