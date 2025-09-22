# Go Genetic Algorithm

This project is a Go-based implementation of a genetic algorithm library. It provides the foundational structures and operations to build and run genetic algorithms for various optimization and search problems. This is currently under development.

I used this project as a learning experience when moving from .NET and C# as my daily driver. I built a similar project that was more specific to one domain in .NET, and decided to create a more generic and reusable library that could be applied to other problem types/domains. 

The project has several comments that were notes to myself for things I ran into while building this, so feel free to ignore those.

## Core Components

*   **Engine:** The central component that orchestrates the genetic algorithm's execution. It manages the population, evolution cycles, and selection process based on the provided configuration.
*   **Configuration:** A flexible `Config` struct allows for tuning the algorithm's parameters, including:
    *   Population Size
    *   Mutation Rate
    *   Crossover Rate
    *   Number of Generations
    *   Elitism (including the number of elites)
    *   Tournament Size (for tournament selection)
    *   Selection Method
    *   Early Stopping (when 100% fitness is reached)
*   **`Individual` Interface:** A generic and well-defined interface for individuals in the population, requiring methods for:
    *   `CalculateFitness()`: Determines the fitness of an individual.
    *   `Crossover()`: Combines two parents to create offspring.
    *   `Mutate()`: Introduces random changes to an individual.
    *   `Fitness()`: Returns the calculated fitness.
    *   `Clone()`: Creates a copy of an individual.
*   **Population Management:** A `Population` struct to manage a collection of individuals.
*   **Selection Strategies:**
    *   Tournament Selection
    *   Roulette Wheel Selection
*   **Callbacks:** A mechanism to monitor the algorithm's progress by setting a fitness callback function.
*   **Working Example:** The main.go file contains a working example for finding a target phrase based on a random input.
    * The configuration can be tweaked to improve fitness. The current example averages around 99.5% using the following config:
      * config := ga.DefaultConfig()
      * config.PopulationSize = 400
      * config.MutationRate = 0.03
      * config.CrossoverRate = 0.9
      * config.Generations = 10000

## Future Work

*   Implement additional selection strategies:
    *   [ ] Rank Selection
*   Add comprehensive unit tests for all components.
*   Add more detailed documentation.

## Areas for Improvement

*   The Tournament Selection works well, but the Roulette Wheel Selection is signifcantly slower and performs worse.
    * This could likely be refined further to optimize this selection method
*   Right now the Fitness() method in the example just calls CalculateFitness() for the sake of the example
    * The actual idea behind it was to use it for caching fitness values to improve performance.
