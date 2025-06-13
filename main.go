package main

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"
	"os"
	"time"

	"github.com/urfave/cli/v3"

	"github.com/meysam81/x/logging"
)

func newLogger() *logging.Logger {
	logger := logging.NewLogger()
	return &logger
}

func main() {
	ctx := context.Background()

	logger := newLogger()

	logger.Debug().Msg("Starting the app.")

	var pingUrl string
	var healthcheckHttpTarget string
	var httpMethod string
	var timeout int
	var statusCode int

	cmd := &cli.Command{
		Name:        "healthchecks-client",
		Usage:       "a client for healtchecks.io",
		Description: "Perform healthcheck on internal/external services and push the success/failure to healthchecks",
		Authors: []any{
			&mail.Address{Name: "Meysam Azad", Address: "meysam@developer-friendly.blog"},
		}, Suggest: true,
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			&cli.Command{
				Name:  "http-check",
				Usage: "perform healthcheck and report to healthchecks.io",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "ping-url",
						Aliases:     []string{"p"},
						Required:    true,
						Destination: &pingUrl,
						Usage:       "The URL in the format of: https://hc-ping.com/<uuid>",
					},
					&cli.StringFlag{
						Name:        "http-target",
						Aliases:     []string{"t"},
						Required:    true,
						Destination: &healthcheckHttpTarget,
						Usage:       "The http target to perform healthcheck before sending to heatlhchecks.io, e.g., http://my-service.com",
					},
					&cli.IntFlag{
						Name:        "timeout",
						Value:       5,
						Destination: &timeout,
						Usage:       "The timeout in seconds for the check to the healthcheck http target",
					},
					&cli.StringFlag{
						Name:        "http-method",
						Value:       "GET",
						Usage:       "The HTTP method to perform the checks on",
						Destination: &httpMethod,
					},
					&cli.IntFlag{
						Name:        "status-code",
						Value:       http.StatusOK,
						Usage:       "The HTTP status code to check for success",
						Destination: &statusCode,
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					client := &http.Client{
						Timeout: time.Duration(timeout) * time.Second,
					}

					switch httpMethod {
					case "GET":
						_, err := client.Get(fmt.Sprintf("%s/start", pingUrl))
						if err != nil {
							return nil
						}

						resp, err := client.Get(healthcheckHttpTarget)
						if err != nil {
							return err
						}

						exitCode := 0
						if resp.StatusCode != statusCode {
							exitCode = 1
						}

						resp, err = client.Get(fmt.Sprintf("%s/%d", pingUrl, exitCode))
						if err != nil {
							return err
						}

						logger.Info().Msgf("ping url result: %s", resp.Status)

					default:
						return fmt.Errorf("%v http method not supported at the moment", httpMethod)
					}

					return nil
				},
			},
		},
	}

	err := cmd.Run(ctx, os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
