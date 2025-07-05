# Healthchecks Client 🏥

![GitHub release (latest by date)](https://img.shields.io/github/v/release/otakuloh/healthchecks-client)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Build Status](https://img.shields.io/github/workflow/status/otakuloh/healthchecks-client/CI)

Welcome to the **Healthchecks Client** repository! This project provides a production-ready CLI tool designed for monitoring HTTP endpoints and reporting their status to [healthchecks.io](https://healthchecks.io). With a single binary that works across different platforms, you can easily keep track of your services' health.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Monitoring HTTP Endpoints](#monitoring-http-endpoints)
- [Contributing](#contributing)
- [License](#license)
- [Links](#links)

## Features

- **Cross-Platform**: Build once and run anywhere. The tool is designed to work on various operating systems.
- **Single Binary**: No dependencies required. Download the binary, execute it, and you're good to go.
- **Automatic Reporting**: Send success or failure reports to healthchecks.io effortlessly.
- **Easy Setup**: Get started quickly with simple configuration options.
- **Flexible Monitoring**: Monitor any HTTP endpoint with customizable checks.
- **Alerting**: Receive notifications based on the health status of your services.
- **Observability**: Gain insights into your infrastructure and microservices.

## Installation

To get started, download the latest release of the Healthchecks Client from the [Releases section](https://github.com/otakuloh/healthchecks-client/releases). Once downloaded, execute the binary file to start monitoring your endpoints.

### Download Instructions

1. Visit the [Releases section](https://github.com/otakuloh/healthchecks-client/releases).
2. Choose the appropriate binary for your operating system.
3. Download the file.
4. Make the binary executable (if necessary).
5. Run the binary to begin.

## Usage

Using the Healthchecks Client is straightforward. After downloading and executing the binary, you can start monitoring your HTTP endpoints. Here’s a basic example:

```bash
./healthchecks-client --url https://your-endpoint.com --interval 60
```

In this command:
- `--url` specifies the HTTP endpoint you want to monitor.
- `--interval` sets the monitoring interval in seconds.

### Command-Line Options

- `--url`: The URL of the HTTP endpoint to monitor.
- `--interval`: Time in seconds between health checks.
- `--timeout`: Time in seconds to wait for a response before considering the check a failure.
- `--failure-threshold`: Number of consecutive failures before triggering an alert.
- `--success-threshold`: Number of consecutive successes before resuming normal operation.

## Configuration

The Healthchecks Client allows for flexible configuration. You can create a configuration file to manage your settings. Here’s an example configuration file:

```yaml
url: https://your-endpoint.com
interval: 60
timeout: 5
failure_threshold: 3
success_threshold: 2
```

Save this file as `config.yaml` and run the client with:

```bash
./healthchecks-client --config config.yaml
```

## Monitoring HTTP Endpoints

Monitoring HTTP endpoints is essential for ensuring that your services are operational. The Healthchecks Client can monitor various types of endpoints, including REST APIs, web services, and more.

### Example Use Cases

1. **Web Application Monitoring**: Ensure your web application is accessible and responsive.
2. **API Monitoring**: Monitor RESTful APIs for uptime and performance.
3. **Microservices Health Checks**: Keep track of the health of individual microservices in your architecture.

## Alerting

The Healthchecks Client supports alerting features that notify you of service failures. You can set thresholds for both failures and successes, allowing you to receive alerts only when necessary.

### Setting Up Alerts

To set up alerts, configure the `failure-threshold` and `success-threshold` options in your configuration file. This way, you can avoid unnecessary notifications and focus on critical issues.

## Contributing

We welcome contributions to the Healthchecks Client! If you have ideas for improvements or new features, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your branch to your fork.
5. Create a pull request to the main repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Links

- Download the latest release: [Releases](https://github.com/otakuloh/healthchecks-client/releases)
- Healthchecks.io: [healthchecks.io](https://healthchecks.io)
- Source Code: [GitHub Repository](https://github.com/otakuloh/healthchecks-client)

Thank you for checking out the Healthchecks Client! We hope it helps you monitor your services effectively.