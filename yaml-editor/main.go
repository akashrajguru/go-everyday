package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// config struct represents the structure of YALM file.
type Config struct {
	Name     string            `yaml:"name"`
	Version  string            `yaml:"version"`
	Database DatabaseConfig    `yaml:"database"`
	Settings map[string]string `yaml:"settings"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
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

	// Lets use writeYAMl function
	// Edit values
	config.Version = "2.0.0"
	config.Database.Host = "localhost"
	config.Database.Port = 5432
	config.Settings["environment"] = "production"
	fmt.Println("Modified config:")
	fmt.Printf("%+v\n\n", config)

	// Write back to file
	err = writeYAML(filename, config)
	if err != nil {
		log.Fatalf("Error writing YAML: %v", err)
	}

	fmt.Println("Successfully updated", filename)
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

func writeYAML(filename string, config *Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failded to marshal YAML: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
