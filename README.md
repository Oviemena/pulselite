# Pulselite
**Tagline**: *Real-time metrics for everyone—cloud, edge, and beyond.*

PulseLite is a lightweight, open-source monitoring tool designed to deliver real-time system metrics with minimal overhead. It consists of two components:
- **Agent**: Collects essential metrics (e.g., CPU, memory, disk, network, uptime, and custom IoT telemetry) and sends them to the aggregator or Prometheus.
- **Aggregator**: Receives metrics, aggregates them, and exposes them via an HTTP API for dashboards or alerting.

Built with Go, PulseLite supports cross-platform compilation for Linux (AMD64 and ARM64), making it ideal for cloud servers, edge devices, and small-scale deployments.

## Vision
PulseLite’s mission is to provide an easy-to-deploy, cost-effective, and customizable monitoring solution for:
- **DevOps Engineers**: Real-time server and container insights with Prometheus compatibility.
- **IoT Developers**: Lightweight metric collection for resource-constrained devices.
- **Small Businesses**: Free, simple monitoring without the complexity or cost of SaaS tools.

## Features
- **Lightweight**: Tiny binary size and low resource usage (~5MB, <1% CPU).
- **Real-Time**: Configurable intervals (as low as 1s) for live metrics.
- **Scalable**: Works on cloud servers, Raspberry Pis, and everything in between.
- **Customizable**: Enable/disable metrics or add custom telemetry (e.g., IoT sensors).
- **Prometheus Compatible**: Export metrics directly to Prometheus for DevOps workflows.
- **Open-Source**: Free forever under the MIT License.

## Metrics
Currently implemented:
- **CPU Usage**: Percentage used across all cores.

Coming soon (see Roadmap):
- Memory usage, disk usage, network I/O, uptime, custom IoT telemetry.

## Prerequisites

To build and run Pulselite, you'll need:
- Go 1.24.1 or later
- Git
- A POSIX-compliant system (Linux, macOS, etc.)

## Building the Binaries

### Clone the Repository

```bash
git clone https://github.com/Oviemena/pulselite.git
cd pulselite
```

### Build from Source

Build the agent and aggregator for your current platform:

```bash
go build -o pulselite-agent cmd/agent/main.go
go build -o pulselite-aggregator cmd/aggregator/main.go
```

Cross-compile for specific architectures:

```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -o pulselite-agent-linux-amd64 cmd/agent/main.go
GOOS=linux GOARCH=amd64 go build -o pulselite-aggregator-linux-amd64 cmd/aggregator/main.go

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -o pulselite-agent-linux-arm64 cmd/agent/main.go
GOOS=linux GOARCH=arm64 go build -o pulselite-aggregator-linux-arm64 cmd/aggregator/main.go
```

## Configuration

Both components use a `config.yaml` file for settings:

```yaml
agent:
  interval: 10s           # How often to collect metrics
  aggregator_addr: "localhost:8080"  # Where to send metrics
  metrics:
    cpu: true              # CPU usage
   

aggregator:
  listen_addr: ":8080"    # HTTP server address
```

## Running the Components

### Start the Aggregator

```bash
./pulselite-aggregator start --config config.yaml
```

Access metrics via HTTP:
```bash
curl http://localhost:8080/stats?name=cpu_usage
```

### Start the Agent

```bash
./pulselite-agent start --config config.yaml
```

## Command-Line Options

Both binaries support:
- `--config <path>`: Specify a custom config file
- `--help`: Show available commands and flags

## CI/CD

GitHub Actions handles:
- Building binaries for Linux (AMD64 and ARM64)
- Creating releases with binary distributions

## Releases

Pre-built binaries are available for Linux (AMD64 and ARM64):
1. Download from the Releases page
2. Extract: `unzip pulselite-binaries-vX.Y.Z.zip`
3. Run as described above

## Contributing

1. Fork the repository
2. Create a feature branch
3. Submit a pull request to main
4. Ensure CI checks pass

## Troubleshooting

- **Agent not sending metrics**: Check `aggregator_addr` in config.yaml
- **Aggregator not responding**: Verify `listen_addr` and port availability
- **Build errors**: Ensure Go 1.24.1 is installed and run `go mod tidy`

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.