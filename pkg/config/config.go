package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Name       string `yaml:"name"`
		UserCol    string `yaml:"userCollectionName"`
		ServiceCol string `yaml:"serviceCollectionName"`
		VersionCol string `yaml:"versionCollectionName"`
		UseMongoDb bool   `yaml:"useMongoDb"`
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
	} `yaml:"database"`
}

// LoadConfig to load configuration from YAML file
func LoadConfig(filePath string) (*Config, error) {
	var config Config
	data, err := readFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("failed to unmarshal yaml file:", err)
		return nil, err
	}
	return &config, nil
}

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("failed to open file:", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("failed to close file:", err)
		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("failed to read file:", err)
		return nil, err
	}
	return data, nil
}
