# PulseLite Configuration Reference

## Configuration File

PulseLite uses YAML for configuration. Default location: `/etc/pulselite/config.yaml`

### Agent Configuration

```yaml
agent:
  # Aggregator endpoint URL
  url: "http://localhost:8080"
  
  # Metric collection interval
  interval: 5s
  
  # Unique identifier for this agent
  source: "my-device"
  
  # Metrics to collect
  metrics:
    cpu_usage: true
    memory_usage: true
    disk_usage: true
    network_io_in: true
    network_io_out: true
    uptime: true
    
  # Debug logging
  verbose: false
```

### Aggregator Configuration

```yaml
aggregator:
  # HTTP API port
  port: "8080"
  
  # Metric retention period
  max_age: 1h
  
  # Debug logging
  verbose: false
```

## Environment Variables

- `PULSELITE_CONFIG`: Override config file path
- `PULSELITE_VERBOSE`: Enable debug logging
- `PULSELITE_PORT`: Override aggregator port