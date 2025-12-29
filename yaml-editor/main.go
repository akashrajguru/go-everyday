package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// config struct represents the structure of YALM file.
type Config struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

func main() {
	filename := "config.yaml"

	// Read file from folder
	config, err := readYAML(filename)
	if err != nil {
		log.Fatalf("Error reading YAML: %v", err)
	}

	fmt.Println("Original config:")
	fmt.Printf("%+v\n\n", config)
}

// Function readYAML to read YAML file and unmarshals it into a config struct
func readYAML(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &config, nil
}
