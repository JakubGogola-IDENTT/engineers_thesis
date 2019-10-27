package config

// Config - configuration of algorithn
type Config struct {
	// NumOfIteration - maximal number of generations
	NumOfIterations uint `json:"NumOfIterations"`

	// SizeOfGenerations - number of speciments in generations
	SizeOfGeneration uint `json:"SizeOfGeneration"`

	// SelectionType - type of selecting speciments to mutation
	SelectionType string `json:"SelectionType"`

	// FromConfig - if set, config should be read from file
	FromFile string

	// PathToImage - path to original image
	PathToImage string `json:"PathToImage"`

	// NumOfBest - number of best speciments used to cross
	NumOfBest uint `json:"NumOfBest"`

	// NumOfThreads - number of threads used to images processing
	NumOfThreads uint `json:"NumOfThreads"`

	// MutationChance - chance of mutation
	MutationChance float64 `json:"MutationChance"`
}
