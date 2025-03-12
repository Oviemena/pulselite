package metrics

import (
    "sync"
    "time"

    "github.com/sirupsen/logrus"
)

type Metric struct {
    Name      string    `json:"name"`
    Value     float64   `json:"value"`
    Timestamp time.Time `json:"timestamp"`
    Source    string    `json:"source"`
}

type MetricStore struct {
    sync.Mutex
    Data   map[string][]Metric
    MaxAge time.Duration
}

func NewMetricStore(maxAge time.Duration) *MetricStore {
    return &MetricStore{
        Data:   make(map[string][]Metric),
        MaxAge: maxAge,
    }
}

func (s *MetricStore) Add(metrics []Metric) {
    s.Lock()
    defer s.Unlock()
    for _, m := range metrics {
        logrus.Debugf("Adding metric: %v", m)
        s.Data[m.Name] = append(s.Data[m.Name], m)
        s.Data[m.Name] = pruneOldMetrics(s.Data[m.Name], s.MaxAge)
        logrus.Debugf("Updated %s slice: %v", m.Name, s.Data[m.Name])
    }
}

func (s *MetricStore) Get(name string) []Metric {
    s.Lock()
    defer s.Unlock()
    return s.Data[name]
}

func pruneOldMetrics(metrics []Metric, maxAge time.Duration) []Metric {
    cutoff := time.Now().UTC().Add(-maxAge)
    var result []Metric
	for _, m := range metrics {
        logrus.Debugf("Checking metric: %v, cutoff: %v, keep: %t", m, cutoff, !m.Timestamp.Before(cutoff))
        if !m.Timestamp.Before(cutoff) { // Keep if not older than cutoff
            result = append(result, m)
        }
    }
    return result
}