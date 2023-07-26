package dummy_connector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

func BuildVersionCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "hub-version",
		Short:   "get version information ",
		Aliases: []string{"version"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return PrintVersion()
		},
	}
	return c
}

func PrintVersion() error {
	fmt.Println("0.0.1")
	return nil
}

func BuildGenerateCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "generate-configuration",
		Short:   "generate-configuration",
		Aliases: []string{"generate"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return GenerateConfiguration()
		},
	}
	return c
}

type exampleConfig struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

var testConfig = exampleConfig{
	Name: "test-config",
	Type: "ndc-config",
}

func GenerateConfiguration() error {

	configJSON, err := json.MarshalIndent(testConfig, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(configJSON))
	return nil
}

func BuildServeCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "serve",
		Short:   "start server",
		Aliases: []string{"start"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return Serve()
		},
	}
	return c
}

func Serve() error {
	fmt.Println("starting server")
	http.HandleFunc("/", rootFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return nil
	}
	return nil
}

func rootFunc(w http.ResponseWriter, req *http.Request) {
	config, err := fetchFileContents("/etc/connector/config.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "file read error"+err.Error())
		fmt.Println("file read error", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, config)
	fmt.Println("responded", config)
}

func fetchFileContents(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read the file: %v", err)
	}
	return string(data), nil
}
