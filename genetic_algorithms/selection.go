package genetic_algorithm

import (
	"errors"
	"math/rand"
)

// TournamentSelection randomly selects `tournamentSize` individuals and returns the best one.
func (p *Population[T]) TournamentSelection(tournamentSize int) (T, error) {
	if len(p.individuals) == 0 {
		return p.handleEmptyIndividuals(p.individuals)
	}

	if tournamentSize > p.size {
		tournamentSize = p.size
	}

	// Select random individuals for the tournament
	tournament := make([]T, tournamentSize)
	for i := 0; i < tournamentSize; i++ {
		tournament[i] = p.individuals[rand.Intn(p.size)]
	}

	// Find the winner
	bestIndividual := tournament[0]
	bestFitness := bestIndividual.CalculateFitness()

	for _, individual := range tournament[1:] {
		fitness := individual.CalculateFitness()
		if fitness > bestFitness {
			bestFitness = fitness
			bestIndividual = individual
		}
	}

	return bestIndividual, nil
}

// RouletteWheelSelection give each individual a chance of being selected proportional to its fitness score.
func (p *Population[T]) RouletteWheelSelection() (T, error) {
	if len(p.individuals) == 0 {
		return p.handleEmptyIndividuals(p.individuals)
	}

	totalFitness := 0.0
	for _, individual := range p.individuals {
		totalFitness += individual.Fitness()
	}

	if totalFitness == 0.0 {
		return p.individuals[rand.Intn(len(p.individuals))], nil
	}

	target := totalFitness * rand.Float64()
	currentSum := 0.0

	for _, individual := range p.individuals {
		currentSum += individual.Fitness()
		if currentSum > target {
			return individual, nil
		}
	}

	// This is a fallback (should never be reached, but here in case of floating point error)
	return p.individuals[rand.Intn(len(p.individuals))], nil
}

// TODO: Implement Rank Selection
//func (p *Population[T]) RankSelection() (T, error) {
//	if len(p.individuals) == 0 {
//		return p.handleEmptyIndividuals(p.individuals)
//	}
//}

// handleEmptyIndividuals is a private method for returning the correct error when p.individuals is empty.
func (p *Population[T]) handleEmptyIndividuals(individuals []T) (T, error) {
	var zeroValue T
	return zeroValue, errors.New(`the "individuals" field of the population is empty`)
}
