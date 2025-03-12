# PulseLite
Lightweight, push-based metrics collector and aggregator.

## Quickstart
1. Build: `go build -o pulselite-agent cmd/agent/main.go && go build -o pulselite-aggregator cmd/aggregator/main.go`
2. Run Aggregator: `./pulselite-aggregator start --port=8080`
3. Run Agent: `./pulselite-agent start --url=http://localhost:8080`
4. Query: `curl http://localhost:8080/stats?name=cpu_usage`

## Features
- Collects CPU usage metrics.
- Stores metrics in memory with configurable retention (default: 1 hour).
- HTTP API for querying.

## License
MIT