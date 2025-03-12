package main

import (
    "bytes"
    "encoding/json"
    "net/http" // Added this import
    "os"
    "time"

    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "github.com/Oviemena/pulselite/pkg/metrics"

    "github.com/shirou/gopsutil/v3/cpu"
)

var (
    aggregatorURL string
    interval      time.Duration
    sourceID      string
    verbose       bool
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "agent",
        Short: "PulseLite agent collects and sends metrics",
    }

    startCmd := &cobra.Command{
        Use:   "start",
        Short: "Start the agent",
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
    cpuPercent, err := cpu.Percent(time.Second, false)
    if err != nil {
        return nil, err
    }
    hostname := sourceID
    if hostname == "" {
        hostname, _ = os.Hostname()
    }
    metric := []metrics.Metric{
        {Name: "cpu_usage", Value: cpuPercent[0], Timestamp: time.Now().UTC(), Source: hostname},
    }
    logrus.Debugf("Collected metrics: %v", metric) // Add this
    return metric, nil
}