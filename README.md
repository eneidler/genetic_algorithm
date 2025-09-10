# Go Genetic Algorithm

This project is a Go-based implementation of a genetic algorithm library. It provides the foundational structures and operations to build and run genetic algorithms for various optimization and search problems. This is a portfolio project and is currently under development.

## Current Features

*   **Generic-based:** The library uses Go generics to allow for flexible and type-safe implementations of `Individual`s.
*   **`Individual` Interface:** A well-defined interface for individuals in the population, including methods for:
    *   `CalculateFitness()`: Determines the fitness of an individual.
    *   `Crossover()`: Combines two parents to create offspring.
    *   `Mutate()`: Introduces random changes to an individual.
    *   `Fitness()`: Returns the calculated fitness if stored, or calls `CalculateFitness()`.
    *   `Clone()`: Creates a copy of an individual.
*   **`Population` Management:** A `Population` struct to manage a collection of individuals.
*   **Selection Strategies:**
    *   Tournament Selection
    *   Roulette Wheel Selection

## Future Work

*   Implement additional selection strategies:
    *   [ ] Rank Selection
*   Implement a complete example in `main.go` to showcase the library's usage.
*   Add more configuration options to the genetic algorithm (e.g., crossover rate, elitism).
*   Write unit tests for all components.
*   Add more comprehensive documentation.
*   Add ga.go file to contain engine logic and config


