# Stage 1: Build the binaries
FROM golang:1.24.1 AS builder

WORKDIR /app

# Copy go.mod and initialize modules
COPY go.mod ./

# Download and verify dependencies
RUN go mod download && \
    go get -d github.com/shirou/gopsutil/v3/cpu@v3.24.5 && \
    go get -d github.com/tklauser/go-sysconf && \
    go mod tidy

# Copy the source code
COPY . .

# Fetch all dependencies explicitly
RUN go get -d gopkg.in/yaml.v3 && \
    go get -d github.com/sirupsen/logrus && \
    go get -d github.com/spf13/cobra


# Fetch missing dependencies
RUN go get -d gopkg.in/yaml.v3 && \
    go get -d github.com/sirupsen/logrus && \
    go get -d github.com/shirou/gopsutil/v3 && \
    go get -d github.com/spf13/cobra


# Build the agent and aggregator binaries for Linux/AMD64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o pulselite-agent-amd64 cmd/agent/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o pulselite-aggregator-amd64 cmd/aggregator/main.go
# Build the agent and aggregator binaries for Linux/ARM64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -o pulselite-agent-arm64 cmd/agent/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -o pulselite-aggregator-arm64 cmd/aggregator/main.go
# Stage 2: Create minimal agent image (AMD64)
FROM alpine:latest AS agent-amd64
COPY --from=builder /app/pulselite-agent-amd64 /usr/local/bin/pulselite-agent
COPY config.yaml /etc/pulselite/config.yaml
ENTRYPOINT ["pulselite-agent", "start"]

# Stage 2: Create minimal agent image (ARM64)
FROM alpine:latest AS agent-arm64
COPY --from=builder /app/pulselite-agent-arm64 /usr/local/bin/pulselite-agent
COPY config.yaml /etc/pulselite/config.yaml
ENTRYPOINT ["pulselite-agent", "start"]

# Stage 2: Create minimal aggregator image (AMD64)
FROM alpine:latest AS aggregator-amd64
COPY --from=builder /app/pulselite-aggregator-amd64 /usr/local/bin/pulselite-aggregator
COPY config.yaml /etc/pulselite/config.yaml
EXPOSE 8080
ENTRYPOINT ["pulselite-aggregator", "start"]

# Stage 2: Create minimal aggregator image (ARM64)
FROM alpine:latest AS aggregator-arm64
COPY --from=builder /app/pulselite-aggregator-arm64 /usr/local/bin/pulselite-aggregator
COPY config.yaml /etc/pulselite/config.yaml
EXPOSE 8080
ENTRYPOINT ["pulselite-aggregator", "start"]