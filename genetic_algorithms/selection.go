package genetic_algorithm

import (
	"math/rand"
)

func (p *Population[T]) TournamentSelection(tournamentSize int) T {
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

	return bestIndividual
}

func (p *Population[T]) RouletteWheelSelection() T {
	// TODO: Implement roulette logic
	var toBeImplemented T
	return toBeImplemented
}

func (p *Population[T]) RankSelection() T {
	// TODO: Implement rank logic
	var toBeImplemented T
	return toBeImplemented
}
