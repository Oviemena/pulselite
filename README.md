# Pulselite

<div align="center">

**Tagline**: *Real-time metrics for everyone—cloud, edge, and beyond.*

[![Go Version](https://img.shields.io/badge/Go-1.24.1-blue.svg)](https://golang.org/doc/go1.24)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

<div>

## Overview


PulseLite is a lightweight, open-source monitoring tool designed to deliver real-time system metrics with minimal overhead. It consists of two components:

### Key Components

- **Agent**: Collects essential metrics (e.g., CPU, memory, disk, network, uptime, and custom IoT telemetry) and sends them to the aggregator or Prometheus.
- **Aggregator**: Receives metrics, aggregates them, and exposes them via an HTTP API for dashboards or alerting.

Built with Go, PulseLite supports cross-platform compilation for Linux (AMD64 and ARM64), making it ideal for cloud servers, edge devices, and small-scale deployments.

## Vision
PulseLite’s mission is to provide an easy-to-deploy, cost-effective, and customizable monitoring solution for:
- **DevOps Engineers**: Real-time server and container insights with Prometheus compatibility.
- **IoT Developers**: Lightweight metric collection for resource-constrained devices.
- **Small Businesses**: Free, simple monitoring without the complexity or cost of SaaS tools.

## Features
- 🪶 **Lightweight**: Tiny binary size and low resource usage (~5MB, <1% CPU).
- ⚡ **Real-Time**: Configurable intervals (as low as 1s) for live metrics.
- 🎡 **Scalable**: Works on cloud servers, Raspberry Pis, and everything in between.
- 🔧 **Customizable**: Enable/disable metrics via config or add custom telemetry (e.g., IoT sensors).
- 📊 **Prometheus Compatible**: Export metrics directly to Prometheus at `/prometheus` for DevOps workflows.
- 🌐 **Open-Source**: Free forever under the MIT License.

## Metrics
PulseLite currently collects:
- **CPU Usage**: Total percentage across all cores.
- **Memory Usage**: Percentage of memory used.
- **Disk Usage**: Percentage used for the root (`/`) filesystem.
- **Network I/O**: Bytes received (`network_io_in`) and sent (`network_io_out`).
- **Uptime**: System uptime in seconds.
- **Custom**: Add your own metrics via config and code.


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
curl http://localhost:8080/prometheus
```

### Start the Agent

```bash
./pulselite-agent start --config config.yaml
```

## Configuration

Create `config.yaml`:

```yaml
agent:
  url: "http://localhost:8080"
  interval: 5s
  source: "my-device"
  metrics:
    cpu_usage: true
    memory_usage: true
    disk_usage: true
    network_io: true
    uptime: true

aggregator:
  port: "8080"
  max_age: 1h
```

## API Endpoints

- `/stats`: JSON metrics
- `/prometheus`: Prometheus format
- `/health`: Service health

## Documentation

- [Installation Guide](docs/installation.md)
- [Configuration Reference](docs/configuration.md)
- [API Documentation](docs/api.md)
- [Contributing Guidelines](CONTRIBUTING.md)

## Command-Line Options

Both binaries support:
- `--config <path>`: Specify a custom config file(both agent and aggregator)
- `--version`: Show version (v0.1.0)
- `--help`: Show available commands and flags

## Prometheus Integration
- Point Prometheus to `http://<aggregator>:8080/prometheus`

## Prometheus Setup
- Download Prometheus: `prometheus.io`
- Configure `prometheus.yaml`:

```yaml
scrape_configs:
  - job_name: 'pulselite'
    static_configs:
      - targets: ['localhost:8080']
```
- Run: `./prometheus --config.file=prometheus.yaml`
- View at `http://localhost:9090`



## CI/CD

GitHub Actions handles:
- Build: Compiles binaries for Linux (AMD64 and ARM64) on every push to main or PR.
- Release: Creates a zip of binaries for tagged releases (e.g., v0.1.0) at Releases.

## Releases

### Latest Release (v0.1.0)

Pre-built binaries are available for:
- Linux AMD64 (`pulselite-agent-linux-amd64`, `pulselite-aggregator-linux-amd64`)
- Linux ARM64 (`pulselite-agent-linux-arm64`, `pulselite-aggregator-linux-arm64`)

#### Installation

```bash
# Download latest release
curl -LO https://github.com/Oviemena/pulselite/releases/latest/download/pulselite-binaries.zip

# Extract binaries
unzip pulselite-binaries.zip

# Make binaries executable
chmod +x pulselite-*

# Optional: Move to system path
sudo mv pulselite-* /usr/local/bin/
```

#### Release History
- **v0.1.0** (2025-03-13)
  - Initial release
  - Basic metric collection
  - Prometheus compatibility
  - AMD64 and ARM64 support

## Roadmap

- IoT Integration: Add temperature with real sensor support (e.g., Raspberry Pi GPIO).
- Enhanced Metrics: Add per-core CPU stats, additional disk mounts.

## Contributing
# Contributing to PulseLite

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.


## Development Requirements

- Go 1.24.1 or later
- Git
- Make (optional)

## Build and Test

```bash
# Build
go build ./cmd/agent
go build ./cmd/aggregator

# Run tests
go test ./...
```

## Commit Messages

Please use conventional commits:

- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation
- `test:` Tests
- `chore:` Maintenance

## Pull Request Process

1. Update documentation
2. Add tests for new features
3. Ensure tests pass
4. Update CHANGELOG.md
5. Request review



## Code Style

- Follow Go standards
- Run `go fmt` before committing
- Use meaningful variable names


**Issues**: File bugs or ideas at [Issues](https://github.com/Oviemena/pulselite/issues).

## Troubleshooting

- **Agent not sending metrics**: Ensure `agent.url` matches `aggregator.port` in `config.yaml` file and verify `aggregator` is running 
- **Aggregator not responding**:“Address already in use” means `8080` is taken. Use `--port 8081` if `8080` is taken or free the port:
```bash
sudo netstat -tulnp | grep 8080
kill -9 <PID>
```
Also check for port availability
- **Build errors**: Ensure Go 1.24.1 is installed and run `go mod tidy`

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details. 
 
