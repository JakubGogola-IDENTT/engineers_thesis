package config

import "flag"

// Config - configuration of algorithn
type Config struct {
	// numOfIteration - maximal number of generations
	numOfIterations uint

	// sizeOfGenerations - number of speciments in generations
	sizeOfGeneration uint

	//operators - genetic operators used in algorithm
	operator string

	// selectionType - type of selecting speciments to mutation
	selectionType string
}

// parseFlags - parses arguments passed to program in console
func (config *Config) parseFlags() {
	flag.UintVar(&config.numOfIterations, "i", 100, "maximal number of generations")
	flag.UintVar(&config.sizeOfGeneration, "g", 15, "size of generation")
	flag.StringVar(&config.operator, "o", CROSSING, "genetic operator")
	flag.StringVar(&config.selectionType, "s", MIXED, "type of selecting speciments to")
}

// Init - initializes config
func (config *Config) Init() {
	config.parseFlags()
}
