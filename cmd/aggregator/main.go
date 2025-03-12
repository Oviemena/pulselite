package main

import (
    "net/http"
    "time"

    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "github.com/Oviemena/pulselite/pkg/api"
    "github.com/Oviemena/pulselite/pkg/metrics"
    "github.com/Oviemena/pulselite/pkg/config"
)

var (
    port    string
    maxAge  time.Duration = time.Hour
    verbose bool
    store   *metrics.MetricStore
	configFile string
)

func main() {
    store = metrics.NewMetricStore(maxAge)

    rootCmd := &cobra.Command{
        Use:   "aggregator",
        Short: "PulseLite aggregator receives and processes metrics",
        Version: "1.0.0",
    }

    startCmd := &cobra.Command{
        Use:   "start",
        Short: "Start the aggregator",
        Run: func(cmd *cobra.Command, args []string) {
			if configFile != "" {
                cfg, err := config.LoadConfig(configFile)
                if err != nil {
                    logrus.Fatalf("Error loading config: %v", err)
                }
                // Override flags with config values if set
                if cfg.Aggregator.Port != "" {
                    port = cfg.Aggregator.Port
                }
                if cfg.Aggregator.MaxAge != 0 {
                    maxAge = cfg.Aggregator.MaxAge
                }
                verbose = cfg.Aggregator.Verbose // No default override needed
            }
            // Initialize store with maxAge after config is loaded
            store = metrics.NewMetricStore(maxAge)
            setupLogger()
            logrus.Infof("Starting PulseLite Aggregator on :%s", port)
            http.HandleFunc("/metrics", api.HandleMetrics(store))
            http.HandleFunc("/stats", api.HandleStats(store))
            logrus.Fatal(http.ListenAndServe(":"+port, nil))
        },
    }
    startCmd.Flags().StringVar(&port, "port", "8080", "Port to listen on")
    startCmd.Flags().DurationVar(&maxAge, "max-age", time.Hour, "Maximum age of stored metrics")
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