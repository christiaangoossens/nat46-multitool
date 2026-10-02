package main

import (
	"flag"
	"log"
	"os"

	"go.yaml.in/yaml/v3"
)

// Config struct. Note: Lowercase fields are ignored by parser.
type config struct {
	// NAT46 Prefix (not yet net.IP)
	RawNAT46Prefix string `yaml:"nat46_prefix"`
}

func main() {
	logger := log.New(os.Stdout, "[NAT46-Multitool] ", log.Ldate|log.Ltime)

	// Get config path from CLI flags
	var configPath string
	flag.StringVar(&configPath, "config", "/etc/nat46-multitool/config.yaml", "Config path")
	flag.Parse()

	// Open the YAML configuration file.
	data, err := os.ReadFile(configPath)
	if err != nil {
		logger.Fatal(err)
	}

	// Application config
	var cfg config

	// Decode YAML content into the config struct.
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		logger.Fatalf("Error decoding YAML: %v", err)
	}

	// Demo!
	logger.Printf("Complete! %s", cfg.RawNAT46Prefix)
}
