package main

import (
    "bytes"
    "encoding/json"
    "net/http" 
    "os"
    "time"

    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
	"github.com/Oviemena/pulselite/pkg/config"
    "github.com/Oviemena/pulselite/pkg/metrics"

    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/host"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/net"
)

var (
    aggregatorURL string
    interval      time.Duration
    sourceID      string
    verbose       bool
	configFile    string
    cfg           *config.Config
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "agent",
        Short: "PulseLite agent collects and sends metrics",
        Version: "1.0.0",
    }

    startCmd := &cobra.Command{
        Use:   "start",
        Short: "Start the agent",
        PreRun: func(cmd *cobra.Command, args []string) {
            var err error
            if configFile != "" {
                cfg, err = config.LoadConfig(configFile)
                if err != nil {
                    logrus.Fatalf("Error loading config: %v", err)
                }
            }
            if cfg == nil {
                cfg = &config.Config{
                    Agent: config.AgentConfig{
                        URL:      "http://localhost:8080",
                        Interval: 5 * time.Second,
                        Source:   os.Getenv("HOSTNAME"),
                        Metrics: map[string]bool{ 
                            "cpu_usage":      true,
                            "memory_usage":   false,
                            "disk_usage":     true,
                            "network_io_in":  true,
                            "network_io_out": true,
                            "uptime":         true,
                    
                        },
                        Verbose:  false,
                    },
                }
            }
            if cfg.Agent.URL != "" {
                aggregatorURL = cfg.Agent.URL
            }
            if cfg.Agent.Interval != 0 {
                interval = cfg.Agent.Interval
            }
            if cfg.Agent.Source != "" {
                sourceID = cfg.Agent.Source
            }
            verbose = cfg.Agent.Verbose
        },
        Run: func(cmd *cobra.Command, args []string) {
            setupLogger()
            logrus.Infof("Starting PulseLite Agent with source ID: %s", sourceID)
            go sendMetrics()
            select {}
        },
    }
    startCmd.Flags().StringVar(&aggregatorURL, "url", "http://localhost:8080", "Aggregator URL")
    startCmd.Flags().DurationVar(&interval, "interval", 5*time.Second, "Collection interval")
    startCmd.Flags().StringVar(&sourceID, "source", os.Getenv("HOSTNAME"), "Source identifier")
    startCmd.Flags().BoolVar(&verbose, "verbose", false, "Enable verbose logging")
    startCmd.Flags().StringVar(&configFile, "config", "/etc/pulselite/config.yaml", "Path to config.yaml file")
    rootCmd.AddCommand(startCmd)

    if err := rootCmd.Execute(); err != nil {
        logrus.Fatal(err)
    }
}

func setupLogger() {
    logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
    if verbose {
        logrus.SetLevel(logrus.DebugLevel)
    } else {
        logrus.SetLevel(logrus.InfoLevel)
    }
}

func sendMetrics() {
    client := &http.Client{Timeout: 10 * time.Second}
    for {
        metrics, err := collectMetrics()
        if err != nil {
            logrus.Errorf("Failed to collect metrics: %v", err)
            time.Sleep(interval)
            continue
        }
        data, _ := json.Marshal(metrics)
        resp, err := client.Post(aggregatorURL+"/metrics", "application/json", bytes.NewBuffer(data))
        if err != nil {
            logrus.Errorf("Failed to send metrics: %v", err)
        } else {
            resp.Body.Close()
            logrus.Debug("Metrics sent successfully")
        }
        time.Sleep(interval)
    }
}

func collectMetrics() ([]metrics.Metric, error) {
    var collected []metrics.Metric
    hostname := sourceID
    if hostname == "" {
        hostname, _ = os.Hostname()
    }
    now := time.Now().UTC()
    if cfg.Agent.Metrics["cpu_usage"] {
        cpuPercent, err := cpu.Percent(time.Second, false)
        if err != nil {
            logrus.Errorf("Failed to collect CPU: %v", err)
        } else {
            collected = append(collected, metrics.Metric{
                Name:      "cpu_usage",
                Value:     cpuPercent[0],
                Timestamp: now,
                Source:    hostname,
            })
        }
    }
    if cfg.Agent.Metrics["memory_usage"] {
        memStats, err := mem.VirtualMemory()
        if err != nil {
            logrus.Errorf("Failed to collect memory: %v", err)
        } else {
            collected = append(collected, metrics.Metric{
                Name:      "memory_usage",
                Value:     memStats.UsedPercent,
                Timestamp: now,
                Source:    hostname,
            })
        }
    }
    if cfg.Agent.Metrics["disk_usage"] {
        diskStats, err := disk.Usage("/")
        if err != nil {
            logrus.Errorf("Failed to collect disk: %v", err)
        } else {
            collected = append(collected, metrics.Metric{
                Name:      "disk_usage",
                Value:     diskStats.UsedPercent,
                Timestamp: now,
                Source:    hostname,
            })
        }
    }
    if cfg.Agent.Metrics["network_io_in"] {
        netStats, err := net.IOCounters(false)
        if err != nil {
            logrus.Errorf("Failed to collect network: %v", err)
        } else {
            collected = append(collected, metrics.Metric{
                Name:      "network_io_in",
                Value:     float64(netStats[0].BytesRecv),
                Timestamp: now,
                Source:    hostname,
            })
        }
    }
    if cfg.Agent.Metrics["network_io_out"] {
        netStats, err := net.IOCounters(false)
        if err != nil {
            logrus.Errorf("Failed to collect network: %v", err)
        } else {
            collected = append(collected, metrics.Metric{
                Name:      "network_io_out",
                Value:     float64(netStats[0].BytesSent),
                Timestamp: now,
                Source:    hostname,
            })
        }
    }
    if cfg.Agent.Metrics["uptime"] {
        uptime, err := host.Uptime()
        if err != nil {
            logrus.Errorf("Failed to collect uptime: %v", err)
        } else {
            collected = append(collected, metrics.Metric{
                Name:      "uptime",
                Value:     float64(uptime),
                Timestamp: now,
                Source:    hostname,
            })
        }
    }
        logrus.Debugf("Collected metrics: %v", collected)
        return collected, nil
    }