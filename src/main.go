package main

import (
	"thesis/config"
	"thesis/genetics"
)

func main() {
	config := &config.Config{}
	config.Init(false)

	genetics.Run("./mona.jpg", g genetics.Genetics)
}
