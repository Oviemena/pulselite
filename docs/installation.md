# PulseLite Installation Guide

## Pre-built Binaries

### Requirements
- Linux (AMD64 or ARM64)
- curl or wget
- unzip

### Steps

1. Download the latest release:
```bash
curl -LO https://github.com/Oviemena/pulselite/releases/latest/download/pulselite-binaries.zip
```

2. Extract the binaries:
```bash
unzip pulselite-binaries.zip
```

3. Make the binaries executable:
```bash
chmod +x pulselite-*
```

4. Optional: Move to system path:
```bash
sudo mv pulselite-* /usr/local/bin/
```

## Building from Source

## Prerequisites

To build and run Pulselite, you'll need:
- **Go**: 1.24.1 or later (for building from source).
- **Git**: To clone the repository.
- **OS**: Linux, macOS, or other POSIX-compliant systems.
- **Make**(optional)


### Steps

1. Clone the repository:
```bash
git clone https://github.com/Oviemena/pulselite.git
cd pulselite
```

2. Build the binaries:
```bash
go build -o pulselite-agent cmd/agent/main.go
go build -o pulselite-aggregator cmd/aggregator/main.go
```

## System Service Setup

### Systemd Service (Linux)

1. Create agent service file:
```bash
sudo nano /etc/systemd/system/pulselite-agent.service
```

2. Add configuration:
```ini
[Unit]
Description=PulseLite Agent
After=network.target

[Service]
ExecStart=/usr/local/bin/pulselite-agent start
Restart=always
User=nobody

[Install]
WantedBy=multi-user.target
```

3. Enable and start service:
```bash
sudo systemctl enable pulselite-agent
sudo systemctl start pulselite-agent
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
