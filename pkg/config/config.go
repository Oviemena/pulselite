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
    Metrics  map[string]bool     `yaml:"metrics"`
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

    if cfg.Agent.URL == "" {
        cfg.Agent.URL = "http://localhost:8080"
    }
    if cfg.Agent.Interval == 0 {
        cfg.Agent.Interval = 5 * time.Second
    }
    if cfg.Agent.Metrics == nil {
        cfg.Agent.Metrics = map[string]bool{
            "cpu_usage":      true,
            "memory_usage":   true,
            "disk_usage":     true,
            "network_io_in":  true,
            "network_io_out": true,
            "uptime":         true,
        }
    }
    if cfg.Aggregator.Port == "" {
        cfg.Aggregator.Port = "8080"
    }
    if cfg.Aggregator.MaxAge == 0 {
        cfg.Aggregator.MaxAge = time.Hour
    }

    return &cfg, nil
}