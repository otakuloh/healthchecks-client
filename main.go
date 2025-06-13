package main

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"
	"os"
	"runtime"
	"time"

	"github.com/urfave/cli/v3"

	"github.com/meysam81/x/logging"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

type HealthcheckConfig struct {
	PingURL    string
	Target     string
	Method     string
	Timeout    time.Duration
	StatusCode int
}

type AppState struct {
	logger *logging.Logger
}

func newAppState() *AppState {
	logger := logging.NewLogger()
	return &AppState{
		logger: &logger,
	}
}

func createHTTPClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}

func sendPingStart(client *http.Client, pingURL string) error {
	_, err := client.Get(fmt.Sprintf("%s/start", pingURL))
	return err
}

func sendPingResult(client *http.Client, pingURL string, exitCode int) (*http.Response, error) {
	return client.Get(fmt.Sprintf("%s/%d", pingURL, exitCode))
}

func (a *AppState) performHealthcheck(client *http.Client, config HealthcheckConfig) (int, error) {
	resp, err := client.Get(config.Target)
	if err != nil {
		return 1, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			a.logger.Warn().Err(closeErr).Msg("Failed to close response body")
		}
	}()

	if resp.StatusCode != config.StatusCode {
		return 1, nil
	}

	return 0, nil
}

func (a *AppState) executeHTTPCheck(ctx context.Context, config HealthcheckConfig) error {
	client := createHTTPClient(config.Timeout)

	if err := sendPingStart(client, config.PingURL); err != nil {
		a.logger.Warn().Err(err).Msg("Failed to send ping start")
	}

	exitCode, err := a.performHealthcheck(client, config)
	if err != nil {
		a.logger.Error().Err(err).Msg("Healthcheck failed")
		exitCode = 1
	}

	resp, err := sendPingResult(client, config.PingURL, exitCode)
	if err != nil {
		return fmt.Errorf("failed to send ping result: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			a.logger.Warn().Err(closeErr).Msg("Failed to close response body")
		}
	}()

	a.logger.Info().Msgf("ping url result: %s", resp.Status)
	return nil
}

func (a *AppState) createHTTPCheckCommand() *cli.Command {
	var config HealthcheckConfig
	var timeoutSeconds int

	return &cli.Command{
		Name:  "http-check",
		Usage: "perform healthcheck and report to healthchecks.io",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "ping-url",
				Aliases:     []string{"p"},
				Required:    true,
				Destination: &config.PingURL,
				Usage:       "The URL in the format of: https://hc-ping.com/<uuid>",
			},
			&cli.StringFlag{
				Name:        "http-target",
				Aliases:     []string{"t"},
				Required:    true,
				Destination: &config.Target,
				Usage:       "The http target to perform healthcheck before sending to heatlhchecks.io, e.g., http://my-service.com",
			},
			&cli.IntFlag{
				Name:        "timeout",
				Value:       5,
				Destination: &timeoutSeconds,
				Usage:       "The timeout in seconds for the check to the healthcheck http target",
			},
			&cli.StringFlag{
				Name:        "http-method",
				Value:       "GET",
				Usage:       "The HTTP method to perform the checks on",
				Destination: &config.Method,
			},
			&cli.IntFlag{
				Name:        "status-code",
				Value:       http.StatusOK,
				Usage:       "The HTTP status code to check for success",
				Destination: &config.StatusCode,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			config.Timeout = time.Duration(timeoutSeconds) * time.Second

			switch config.Method {
			case "GET":
				return a.executeHTTPCheck(ctx, config)
			default:
				return fmt.Errorf("%s http method not supported at the moment", config.Method)
			}
		},
	}
}

func createVersionCommand() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "show version information",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Printf("Version:    %s\n", version)
			fmt.Printf("Commit:     %s\n", commit)
			fmt.Printf("Built:      %s\n", date)
			fmt.Printf("Built by:   %s\n", builtBy)
			fmt.Printf("Go version: %s\n", runtime.Version())
			fmt.Printf("OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
			return nil
		},
	}
}

func (a *AppState) createRootCommand() *cli.Command {
	return &cli.Command{
		Name:        "healthchecks-client",
		Usage:       "a client for healtchecks.io",
		Description: "Perform healthcheck on internal/external services and push the success/failure to healthchecks",
		Version:     version,
		Authors: []any{
			&mail.Address{Name: "Meysam Azad", Address: "meysam@developer-friendly.blog"},
		},
		Suggest:               true,
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			a.createHTTPCheckCommand(),
			createVersionCommand(),
		},
	}
}

func main() {
	ctx := context.Background()
	app := newAppState()
	app.logger.Debug().Msg("Starting the app.")

	cmd := app.createRootCommand()
	if err := cmd.Run(ctx, os.Args); err != nil {
		app.logger.Fatal().Err(err).Msg("Application failed")
	}
}
