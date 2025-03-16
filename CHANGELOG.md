# Changelog

All notable changes to **PulseLite** will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
- No unreleased changes yet.

## [0.2.0] - 2025-03-14
### Added
- Windows OS support for agent and aggregator (metrics adapted for `C:` disk usage, previously relied on `gopsutil` fallback).


## [0.2.0] -2025-03-14
### Changed
- Updated `disk_usage` to use `C:` on Windows and `/` on POSIX systems.
- Enhanced `README.md` with Windows build/run instructions and setup.
- Improved metric readability: `cpu_usage`, `memory_usage`, `disk_usage` rounded to 2 decimals.
- Optimised `release.yml` and `ci.ym` files.

### Fixed
- Config parsing errors from `[]string` to `map[string]bool` mismatch in earlier iterations.

## [0.1.0] - 2025-03-13
### Added
- Initial release of PulseLite with agent and aggregator components.
- Core metrics: `cpu_usage`, `memory_usage`, `disk_usage`, `network_io_in`, `network_io_out`, `uptime`.
- Prometheus exporter at `/prometheus` endpoint.
- Configurable metrics via `config.yaml` (`metrics: map[string]bool`).
- HTTP API endpoints: `/stats` for querying metrics, `/metrics` for ingestion.
- Cross-platform builds for Linux (AMD64, ARM64).
- MIT License and basic `README.md` documentation.

---

[Unreleased]: https://github.com/Oviemena/pulselite/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/Oviemena/pulselite/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/Oviemena/pulselite/releases/tag/v0.1.0