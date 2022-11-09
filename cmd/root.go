/*
Copyright © 2019 NAME HERE <hello@lukasbahr.de>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	// "context"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	deconz_api_key string
	deconz_api_url string
	deconz_device  string
	log_level      string
	buildVersion   string
	buildCommit    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "deconz-ctl",
	Short: "cli client for deconz",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey, baseurl string) *Client {
	return &Client{
		BaseURL: baseurl,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLogger)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("deconz")
	rootCmd.PersistentFlags().StringVarP(&deconz_device, "device", "d", viper.GetString("device"), "the device to handle")
	rootCmd.PersistentFlags().StringVarP(&deconz_api_url, "apiUrl", "u", viper.GetString("api_url"), "the deconz api address")
	rootCmd.PersistentFlags().StringVarP(&deconz_api_key, "apiKey", "k", viper.GetString("api_key"), "the deconz api key")
	rootCmd.PersistentFlags().StringVarP(&log_level, "loglevel", "l", "info", "the loglevel")
	rootCmd.AddCommand(switchOnCmd)
	rootCmd.AddCommand(switchOffCmd)
	rootCmd.AddCommand(versionCmd)
}

func initLogger() {

	var l log.Level

	switch log_level {
	case "panic":
		l = log.PanicLevel
	case "fatal":
		l = log.FatalLevel
	case "error":
		l = log.ErrorLevel
	case "warn", "warning":
		l = log.WarnLevel
	case "info":
		l = log.InfoLevel
	case "debug":
		l = log.DebugLevel
	case "trace":
		l = log.TraceLevel
	}

	formatter := &logrus.TextFormatter{DisableQuote: true}
	formatter.DisableQuote = true

	log.SetFormatter(formatter)
	log.SetLevel(l)

	log.SetOutput(os.Stdout)
}

var switchOnCmd = &cobra.Command{
	Use:   "switch-on",
	Short: "switch-on allows to switch off a light",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		body, err := sendRequest(ctx, true)
		log.Debugf("switching on device %s on %s", deconz_device, deconz_api_url)
		if err != nil {
			log.Errorf("error sending request to deconz appliance %s", err.Error())
		}
		log.Debugf("deconz responded with body: %s", body)
	},
}

var switchOffCmd = &cobra.Command{
	Use:   "switch-off",
	Short: "switch-off allows to switch off a light",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		body, err := sendRequest(ctx, false)
		log.Debugf("switching on device %s on %s", deconz_device, deconz_api_url)
		if err != nil {
			log.Errorf("error sending request to deconz appliance %s", err.Error())
		}
		log.Debugf("deconz responded with body: %s", body)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version shows the version of the cli tool",
	Long:  `All software has versions. This is deconz-ctl's`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("version %s commit sha %s", buildVersion, buildCommit)
	},
}

func sendRequest(ctx context.Context, state bool) ([]byte, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"on": state,
	})
	if err != nil {
		log.Fatal(err)
	}

	if deconz_device == "" {
		log.Fatalf("deconz device to handle is not set. Please specify device by using -d parameter or DECONZ_DEVICE environment variable ")
	}
	request_url := fmt.Sprintf("%s/api/%s/lights/%s/state", deconz_api_url, deconz_api_key, deconz_device)
	req, err := http.NewRequestWithContext(ctx, "PUT", request_url, bytes.NewBuffer(payload))
	if err != nil {
		log.Errorf("error setting up request %s", err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("error setting up request %s", err.Error())
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Errorf("error reading response body %s", err.Error())
		return nil, err
	}
	return body, nil
}
