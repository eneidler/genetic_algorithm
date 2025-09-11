package ga

import (
	"fmt"
	"math/rand"
	"time"
)

type Engine[T Individual] struct {
	// Core params
	rng            *rand.Rand
	populationSize int
	mutationRate   float64
	crossoverRate  float64
	generations    int

	// Advanced params
	elitism         bool
	eliteCount      int
	tournamentSize  int
	selectionMethod SelectionMethod

	// functions
	createIndividual func() T
	fitnessCallback  func(T, float64, int)

	// State
	population     *Population[T]
	bestIndividual T
	bestFitness    float64
	generation     int
	fitnessHistory []float64
}

// CreateEngine returns a pointer to a new Engine, as well as standard error handling.
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

	// Handles elitism logic if it's configured and enabled
	if config.Elitism {
		if config.EliteCount <= 0 {
			config.EliteCount = config.PopulationSize / 10 // Ensure at least 10% minimum elitism
			if config.EliteCount < 1 {
				config.EliteCount = 1
			}
		} else if config.EliteCount > config.PopulationSize/2 { // limit elitism to 50% max
			config.EliteCount = config.PopulationSize / 2
			if config.EliteCount < 1 {
				config.EliteCount = 1
			}
		}
	}

	ga := &Engine[T]{
		rng:              rand.New(rand.NewSource(time.Now().UnixNano())),
		populationSize:   config.PopulationSize,
		mutationRate:     config.MutationRate,
		crossoverRate:    config.CrossoverRate,
		generations:      config.Generations,
		elitism:          config.Elitism,
		eliteCount:       config.EliteCount,
		tournamentSize:   config.TournamentSize,
		selectionMethod:  config.SelectionMethod,
		createIndividual: createIndividual,
		fitnessHistory:   make([]float64, 0, config.Generations),
	}

	return ga, nil
}

func (e *Engine[T]) Run() (T, float64, error) {

}
