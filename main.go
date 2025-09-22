package main

import (
	"fmt"
	"math/rand"

	ga "github.com/eneidler/genetic_algorithms/genetic_algorithms"
)

// StringIndividual represents a solution as a string
type StringIndividual struct {
	genes  string
	target string
}

func (s *StringIndividual) CalculateFitness() float64 {
	fitness := 0.0
	for i := 0; i < len(s.genes) && i < len(s.target); i++ {
		if s.genes[i] == s.target[i] {
			fitness += 1.0
		}
	}
	return fitness / float64(len(s.target)) * 100 // Percentage match
}

func (s *StringIndividual) Crossover(other ga.Individual) []ga.Individual {
	// Self-note: This type assertion is needed because `other` comes in as generic ga.Individual.
	// This converts it to the concrete type. This tripped me up the first time around while trying to learn Go
	// through this project when `genes` gave me an unresolved reference.
	otherString := other.(*StringIndividual)

	// Single-point crossover
	crossoverPoint := rand.Intn(len(s.genes))

	child1Genes := s.genes[:crossoverPoint] + otherString.genes[crossoverPoint:]
	child2Genes := otherString.genes[:crossoverPoint] + s.genes[crossoverPoint:]

	return []ga.Individual{
		&StringIndividual{genes: child1Genes, target: s.target},
		&StringIndividual{genes: child2Genes, target: s.target},
	}
}

func (s *StringIndividual) Mutate(crossoverRate float64) {
	// Random mutation of one character
	if len(s.genes) == 0 {
		return
	}

	pos := rand.Intn(len(s.genes))
	chars := []rune(s.genes)
	chars[pos] = rune(rand.Intn(94) + 32) // Random printable ASCII
	s.genes = string(chars)
}

func (s *StringIndividual) Fitness() float64 {
	return s.CalculateFitness()
}

func (s *StringIndividual) Clone() ga.Individual {
	return &StringIndividual{
		genes:  s.genes,
		target: s.target,
	}
}

func main() {
	// I stated with a basic "Hello World" here, but I kept pushing the complexity of target to see
	// how it performed.
	target := "It was the best of times, it was the worst of times, " +
		"it was the age of wisdom, it was the age of foolishness, it was the epoch of belief, " +
		"it was the epoch of incredulity, it was the season of light, it was the season of darkness, " +
		"it was the spring of hope, it was the winter of despair"

	createIndividual := func() *StringIndividual {
		// Generate random string of same length as target
		genes := make([]byte, len(target))
		for i := range genes {
			genes[i] = byte(rand.Intn(94) + 32) // Random printable ASCII
		}
		return &StringIndividual{
			genes:  string(genes),
			target: target,
		}
	}

	config := ga.DefaultConfig()
	config.PopulationSize = 400
	config.MutationRate = 0.03
	config.CrossoverRate = 0.9
	config.Generations = 20000
	config.EarlyStopping = true
	// config.Elitism = false
	config.EliteCount = 4
	// config.TournamentSize = 2
	// config.SelectionMethod = ga.RouletteWheelSelection

	engine, err := ga.CreateEngine(config, createIndividual)
	if err != nil {
		fmt.Println(err)
	}

	// Add callback to see progress
	engine.SetFitnessCallback(func(best *StringIndividual, fitness float64, gen int) {
		if gen%100 == 0 {
			fmt.Printf("Gen %d: %s (fitness: %.2f%%)\n", gen, best.genes, fitness)
		}
	})

	best, fitness, gen, err := engine.Run()
	fmt.Printf("Final result: %s (fitness: %.2f%%, final gen: %d)\n", best.genes, fitness, gen)
}
