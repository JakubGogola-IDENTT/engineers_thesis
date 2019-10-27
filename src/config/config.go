package config

import (
	"flag"
	"fmt"
)

// Config - configuration of algorithn
type Config struct {
	// NumOfIteration - maximal number of generations
	NumOfIterations uint

	// SizeOfGenerations - number of speciments in generations
	SizeOfGeneration uint

	// Operators - genetic operators used in algorithm
	Operator string

	// SelectionType - type of selecting speciments to mutation
	SelectionType string

	// FromConfig - if set, config should be read from file
	FromFile bool

	// PathToImage - path to original image
	PathToImage string

	// NumOfBest - number of best speciments used to cross
	NumOfBest uint

	// NumOfThreads - number of threads used to images processing
	NumOfThreads uint
}

// parseFlags - parses arguments passed to program in console
func (config *Config) parseFlags() {
	flag.BoolVar(&config.FromFile, "f", false, "read config from file")
	flag.UintVar(&config.NumOfIterations, "i", 100, "maximal number of generations")
	flag.UintVar(&config.SizeOfGeneration, "g", 300, "size of generation")
	flag.UintVar(&config.NumOfBest, "b", 2, "number of best specimens")
	flag.UintVar(&config.NumOfThreads, "t", 1, "number of threads")
	flag.StringVar(&config.Operator, "o", CROSSING, "genetic operator")
	flag.StringVar(&config.SelectionType, "s", MIXED, "type of selecting speciments to")
	flag.StringVar(&config.PathToImage, "d", "./mona.jpg", "path to original image")
}

// Init - initializes config
func (config *Config) Init(fromFile bool) { // TODO: add reading from file guard
	if fromFile {
		fmt.Println("Read config from file")
	}

	config.parseFlags()
	flag.Parse()
}
