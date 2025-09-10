package ga

// SelectionMethod defines the type of selection method to use
type SelectionMethod int

const (
	TournamentSelection SelectionMethod = iota
	RouletteWheelSelection
)

type Config struct {
	PopulationSize  int
	MutationRate    float64
	CrossoverRate   float64
	Generations     int
	Elitism         bool
	EliteCount      int
	TournamentSize  int
	SelectionMethod SelectionMethod
}

func DefaultConfig() Config {
	return Config{
		PopulationSize:  100,
		MutationRate:    0.01,
		CrossoverRate:   0.8,
		Generations:     1000,
		Elitism:         true,
		EliteCount:      2,
		TournamentSize:  5,
		SelectionMethod: TournamentSelection,
	}
}
