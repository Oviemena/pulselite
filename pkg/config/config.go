package config

import (
    "fmt"
    "os"
    "time"

    "gopkg.in/yaml.v3"
)

type Config struct {
    Agent      AgentConfig      `yaml:"agent"`
    Aggregator AggregatorConfig `yaml:"aggregator"`
}

type AgentConfig struct {
    URL      string        `yaml:"url"`
    Interval time.Duration `yaml:"interval"`
    Source   string        `yaml:"source"`
    Metrics  []string      `yaml:"metrics"`
    Verbose  bool          `yaml:"verbose"`
}

type AggregatorConfig struct {
    Port    string        `yaml:"port"`
    MaxAge  time.Duration `yaml:"max_age"`
    Verbose bool          `yaml:"verbose"`
}

func LoadConfig(file string) (*Config, error) {
    data, err := os.ReadFile(file)
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %v", err)
    }

    var cfg Config
    if err := yaml.Unmarshal(data, &cfg); err != nil {
        return nil, fmt.Errorf("failed to parse config file: %v", err)
    }

    // Set defaults if fields are omitted
    if cfg.Agent.URL == "" {
        cfg.Agent.URL = "http://localhost:8080"
    }
    if cfg.Agent.Interval == 0 {
        cfg.Agent.Interval = 5 * time.Second
    }
    if len(cfg.Agent.Metrics) == 0 {
        cfg.Agent.Metrics = []string{"cpu_usage"}
    }
    if cfg.Aggregator.Port == "" {
        cfg.Aggregator.Port = "8080"
    }
    if cfg.Aggregator.MaxAge == 0 {
        cfg.Aggregator.MaxAge = time.Hour
    }

    return &cfg, nil
}