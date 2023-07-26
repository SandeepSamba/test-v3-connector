package main

import (
	"fmt"
	"os"

	"github.com/hasura/dummy_connector"
	"github.com/spf13/cobra"
)

func main() {
	if err := execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func execute() error {
	var err error

	cmd := buildRootCommand()
	cmd, err = cmd.ExecuteC()

	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func buildRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "connector",
		Short:         "Hasura Connector CLI.",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	// cmd.PersistentFlags().StringVar(&hasuraConnector.LogLevel, "log-level", "", "enable debug logging")

	cmd.AddCommand(dummy_connector.BuildVersionCommand())
	cmd.AddCommand(dummy_connector.BuildGenerateCommand())
	cmd.AddCommand(dummy_connector.BuildServeCommand())

	return cmd
}
