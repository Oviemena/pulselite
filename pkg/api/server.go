package api

import (
    "encoding/json"
    "net/http"

    "github.com/sirupsen/logrus"
    "github.com/Oviemena/pulselite/pkg/metrics"
)

func HandleMetrics(store *metrics.MetricStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var metrics []metrics.Metric
        if err := json.NewDecoder(r.Body).Decode(&metrics); err != nil {
            logrus.Errorf("Failed to decode metrics: %v", err)
            http.Error(w, "Invalid payload", http.StatusBadRequest)
            return
        }
        logrus.Debugf("Received %d metrics: %v", len(metrics), metrics)
        store.Add(metrics) // Run synchronously for now
        logrus.Debugf("Stored %d metrics, current store: %v", len(metrics), store.Data)
        w.Write([]byte("OK"))
    }
}

func HandleStats(store *metrics.MetricStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        metricName := r.URL.Query().Get("name")
        store.Lock()
        defer store.Unlock()
        logrus.Debugf("Query for metric: %s", metricName)
        logrus.Debugf("Current store data: %v", store.Data)
        if metricName == "" {
            logrus.Debug("Returning all metrics")
            json.NewEncoder(w).Encode(store.Data)
            return
        }
        if data, ok := store.Data[metricName]; ok {
            logrus.Debugf("Found data for %s: %v", metricName, data)
            w.Header().Set("Content-Type", "application/json")
            jsonBytes, err := json.Marshal(data)
            if err != nil {
                logrus.Errorf("Failed to marshal data: %v", err)
                http.Error(w, "Internal server error", http.StatusInternalServerError)
                return
            }
            w.Write(jsonBytes)
            return
        }
        logrus.Debugf("Metric %s not found", metricName)
        http.Error(w, "Metric not found", http.StatusNotFound)
    }
}