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
	earlyStopping  bool
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
		earlyStopping:    config.EarlyStopping,
	}

	return ga, nil
}

func (ga *Engine[T]) SetFitnessCallback(callback func(T, float64, int)) {
	ga.fitnessCallback = callback
}

// Run executes the genetic algorithm functionality and returns the best individual and its fitness
func (ga *Engine[T]) Run() (T, float64, int, error) {
	ga.population = NewPopulation(ga.populationSize, ga.createIndividual)

	// Track the best individual across all generations
	ga.bestIndividual, ga.bestFitness = ga.population.GetBest()

	var finalGen int

	// Evolution loop
	for gen := 0; gen < ga.generations; gen++ {
		ga.generation = gen

		nextGen, err := ga.evolvePopulation()
		if err != nil {
			return ga.bestIndividual, ga.bestFitness, gen, err
		}

		ga.population = nextGen

		currentBest, currentBestFitness := ga.population.GetBest()
		if currentBestFitness > ga.bestFitness {
			ga.bestIndividual = currentBest
			ga.bestFitness = currentBestFitness
		}

		ga.fitnessHistory = append(ga.fitnessHistory, currentBestFitness)

		if ga.fitnessCallback != nil {
			ga.fitnessCallback(currentBest, currentBestFitness, gen)
		}

		if ga.bestFitness == 100.0 && ga.earlyStopping {
			finalGen = gen
			return ga.bestIndividual, ga.bestFitness, finalGen, nil
		}
	}

	return ga.bestIndividual, ga.bestFitness, finalGen, nil
}

func (ga *Engine[T]) selectParent() (T, error) {
	switch ga.selectionMethod {
	case TournamentSelection:
		return ga.population.TournamentSelection(ga.tournamentSize)
	case RouletteWheelSelection:
		return ga.population.RouletteWheelSelection()
	default:
		return ga.population.TournamentSelection(ga.tournamentSize)
	}
}

func (ga *Engine[T]) evolvePopulation() (*Population[T], error) {
	newIndividuals := make([]T, 0, ga.populationSize)

	// Elitism logic
	if ga.elitism && ga.eliteCount > 0 {
		elite := ga.population.GetElite(ga.eliteCount)
		for _, e := range elite {
			newIndividuals = append(newIndividuals, e.Clone().(T))
		}
	}

	// Crossover and Mutation Logic
	for len(newIndividuals) < ga.populationSize {
		parent1, err := ga.selectParent()
		if err != nil {
			return nil, err
		}

		if rand.Float64() < ga.crossoverRate {
			parent2, err := ga.selectParent()
			if err != nil {
				return nil, err
			}

			offspring := parent1.Crossover(parent2)

			for _, child := range offspring {
				if len(newIndividuals) < ga.populationSize {
					childT := child.(T)

					if rand.Float64() < ga.mutationRate {
						childT.Mutate(ga.mutationRate)
					}

					newIndividuals = append(newIndividuals, childT)
				}
			}
		} else {
			// No crossover, just copy parent with possible mutation
			child := parent1.Clone().(T)

			if rand.Float64() < ga.mutationRate {
				child.Mutate(ga.mutationRate)
			}

			newIndividuals = append(newIndividuals, child)
		}
	}

	newPop := &Population[T]{
		individuals: newIndividuals[:ga.populationSize], // Ensure exact population size
		size:        ga.populationSize,
	}

	return newPop, nil
}
