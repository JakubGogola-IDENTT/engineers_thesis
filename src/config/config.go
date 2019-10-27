package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// parseFlags - parses arguments passed to program in console
func (c *Config) parseFlags() {
	flag.StringVar(&c.FromFile, "from-file", "", "read config from given file")
	flag.UintVar(&c.NumOfIterations, "generations", 100, "maximal number of generations")
	flag.UintVar(&c.SizeOfGeneration, "generation-size", 300, "size of generation")
	flag.UintVar(&c.NumOfBest, "best", 10, "number of best specimens")
	flag.UintVar(&c.NumOfThreads, "threads", 10, "number of threads")
	flag.StringVar(&c.SelectionType, "selection", MIXED, "type of selecting speciments to")
	flag.StringVar(&c.PathToImage, "image-dir", "./mona.jpg", "path to original image")
	flag.Float64Var(&c.MutationChance, "mutation-chance", 0.2, "chance of mutation")
}

// Init - initializes config
func (c *Config) Init(fromFile bool) { // TODO: add reading from file guard
	if fromFile {
		fmt.Println("Read config from file")
	}

	c.parseFlags()
	flag.Parse()

	// read from file if path is not empty
	if c.FromFile != "" {
		c.readFromFile()
	}
}

func (c *Config) readFromFile() {
	jsonFile, err := os.Open(c.FromFile)

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	bytesValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(bytesValue, &c)
}
