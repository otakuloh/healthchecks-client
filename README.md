# healthchecks-client

<div align="center">

<!-- Project Status & Quality -->

[![CI/CD](https://github.com/meysam81/healthchecks-client/actions/workflows/ci.yml/badge.svg)](https://github.com/meysam81/healthchecks-client/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/meysam81/healthchecks-client)](https://goreportcard.com/report/github.com/meysam81/healthchecks-client)
[![Vulnerability Scan](https://img.shields.io/badge/🛡️_Security-Scanned-brightgreen?style=flat-square)](https://github.com/meysam81/healthchecks-client/actions)

<!-- Release & Distribution -->

[![Latest Release](https://img.shields.io/github/v/release/meysam81/healthchecks-client?style=flat-square&logo=github&color=blue)](https://github.com/meysam81/healthchecks-client/releases/latest)
[![Docker Image](https://img.shields.io/badge/docker-meysam81%2Fhealthchecks--client-blue?style=flat-square&logo=docker)](https://ghcr.io/meysam81/healthchecks-client)
[![Go Version](https://img.shields.io/github/go-mod/go-version/meysam81/healthchecks-client?style=flat-square&logo=go)](go.mod)

<!-- License & Community -->

[![License](https://img.shields.io/badge/License-Apache--2.0-green.svg?style=flat-square)](LICENSE)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/meysam81?style=flat-square&logo=github&color=pink)](https://github.com/sponsors/meysam81)

<!-- Technical Features -->

[![Single Binary](https://img.shields.io/badge/🚀_Single-Binary-blueviolet?style=flat-square)](https://golang.org/)
[![Cross Platform](https://img.shields.io/badge/🌐_Cross-Platform-orange?style=flat-square)](https://golang.org/)

<!-- Monitoring Features -->

[![Healthchecks.io](https://img.shields.io/badge/✅_healthchecks.io-Compatible-4CAF50?style=flat-square)](https://healthchecks.io)
[![HTTP Monitoring](https://img.shields.io/badge/🌐_HTTP-Monitoring-2196F3?style=flat-square)](https://developer.mozilla.org/en-US/docs/Web/HTTP)
[![Production Ready](https://img.shields.io/badge/🏭_Production-Ready-darkgreen?style=flat-square)](https://sre.google/)

</div>

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [What it does](#what-it-does)
- [Installation](#installation)
- [Usage](#usage)
  - [HTTP Health Check](#http-health-check)
  - [Options](#options)
  - [Version](#version)
- [Example](#example)
- [How it works](#how-it-works)
- [License](#license)
- [Author](#author)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

A lightweight CLI tool for monitoring services and reporting to [healthchecks.io](https://healthchecks.io).

## What it does

Performs health checks on your services and automatically reports success/failure to healthchecks.io monitoring service.

## Installation

```bash
go install github.com/meysam81/healthchecks-client@latest
```

## Usage

### HTTP Health Check

```bash
healthchecks-client http-check \
  --ping-url https://hc-ping.com/your-uuid \
  --http-target http://your-service.com/health
```

### Options

```
--ping-url, -p      Your healthchecks.io ping URL (required)
--http-target, -t   Service endpoint to check (required)
--timeout           Request timeout in seconds (default: 5)
--http-method       HTTP method to use (default: GET)
--status-code       Expected HTTP status code (default: 200)
```

### Version

```bash
healthchecks-client version
```

## Example

```bash
# Monitor a web service
healthchecks-client http-check \
  -p https://hc-ping.com/12345678-1234-1234-1234-123456789abc \
  -t https://api.example.com/health \
  --timeout 10 \
  --status-code 200
```

## How it works

1. Sends a "start" ping to healthchecks.io
2. Performs the health check on your target service
3. Reports success (exit code 0) or failure (exit code 1) back to healthchecks.io

## License

Apache License 2.0 - see [LICENSE](LICENSE) for details.

## Author

[Meysam Azad](mailto:meysam@developer-friendly.blog)
