package main

import (
	"thesis/config"
)

func main() {
	config := &config.Config{}
	config.Init()
}
