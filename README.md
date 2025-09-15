# Go Genetic Algorithm

This project is a Go-based implementation of a genetic algorithm library. It provides the foundational structures and operations to build and run genetic algorithms for various optimization and search problems. This project is currently under development.

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

## Future Work

*   Implement additional selection strategies:
    *   [ ] Rank Selection
*   Implement a complete example in `main.go` to showcase the library's usage.
*   Add comprehensive unit tests for all components.
*   Add more detailed documentation.
