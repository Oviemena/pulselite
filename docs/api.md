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
unzip pulselite-binaries.zip# PulseLite API Documentation

## Endpoints

### GET /stats
Returns metrics in JSON format.

#### Parameters
- `name` (optional): Filter by metric name
- `source` (optional): Filter by source
- `from` (optional): Start time (ISO 8601)
- `to` (optional): End time (ISO 8601)

#### Example
```bash
curl "http://localhost:8080/stats?name=cpu_usage&source=my-device"
```

#### Response
```json
{
  "cpu_usage": [
    {
      "name": "cpu_usage",
      "value": 1.5,
      "timestamp": "2025-03-13T13:00:00Z",
      "source": "my-device"
    }
  ]
}
```

### GET /prometheus
Returns metrics in Prometheus format.

#### Example
```bash
curl http://localhost:8080/prometheus
```

#### Response
```text
pulselite_cpu_usage{source="my-device"} 1.5
pulselite_memory_usage{source="my-device"} 41.2
```

### GET /health
Returns service health status.

#### Response
```json
{
  "status": "healthy",
  "uptime": "1h2m3s",
  "version": "v0.1.0"
}
```

## Error Responses

```json
{
  "error": "Invalid parameter: name",
  "code": 400
}
```
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

### Requirements
- Go 1.24.1 or later
- Git
- Make (optional)

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