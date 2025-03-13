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
PulseLite currently collects:
- **CPU Usage**: Total percentage across all cores.
- **Memory Usage**: Percentage of memory used.
- **Disk Usage**: Percentage used for the root (`/`) filesystem.
- **Network I/O**: Bytes received (`network_io_in`) and sent (`network_io_out`).
- **Uptime**: System uptime in seconds.

Coming soon (see Roadmap):
-  custom IoT telemetry.

## Prerequisites

To build and run Pulselite, you'll need:
- **Go**: 1.24.1 or later (for building from source).
- **Git**: To clone the repository.
- **OS**: Linux, macOS, or other POSIX-compliant systems.

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
  url: "http://localhost:8080"  # Aggregator endpoint
  interval: 5s                 # Collection interval
  source: "my-device-agent"          # Unique identifier (e.g., hostname)
  metrics:                     # Metrics to collect
    - cpu_usage
    - memory_usage
    - disk_usage
    - network_io_in
    - network_io_out
    - uptime
    - temperature
  verbose: false               # Enable debug logging
aggregator:
  port: "8080"                 # HTTP API port
  max_age: 1h                  # How long to retain metrics
  verbose: false               # Enable debug logging
```

## Running the Components

### Start the Aggregator

```bash
./pulselite-aggregator start --config config.yaml
```
Or without config (uses defaults: port 8080, 1h max age):

```bash
./pulselite-aggregator start
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
- `--config <path>`: Specify a custom config file(both agent and aggregator)
- `--help`: Show available commands and flags

## CI/CD

GitHub Actions handles:
- Build: Compiles binaries for Linux (AMD64 and ARM64) on every push to main or PR.
- Release: Creates a zip of binaries for tagged releases (e.g., v0.1.0) at Releases.

## Releases

Pre-built binaries are available for Linux (AMD64 and ARM64):
1. Download from the Releases page
2. Extract: `unzip pulselite-binaries-vX.Y.Z.zip`
3. Run as described above


## Roadmap

- Customizable Metrics: Enable/disable specific metrics via config with finer control.
- Prometheus Compatibility: Export metrics to Prometheus for DevOps workflows.
- IoT Integration: Replace temperature placeholder with real sensor support (e.g., Raspberry Pi GPIO).
- Enhanced Metrics: Add per-core CPU stats, additional disk mounts.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Submit a pull request to main
4. Ensure CI checks pass

## Troubleshooting

- **Agent not sending metrics**: Ensure `agent.url` matches `aggregator.port` in `config.yaml` file and verify aggregator is running 
- **Aggregator not responding**:“Address already in use” means `8080` is taken. Use `--port 8081` or free the port:
```bash
sudo netstat -tulnp | grep 8080
kill -9 <PID>
```
Also check for port availability
- **Build errors**: Ensure Go 1.24.1 is installed and run `go mod tidy`

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.