package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

type Input struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func Query(input Input) (res *http.Response, err error) {
	jsonValue, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	apiUrl := os.Getenv("RUNPOD_API_URL")
	if apiUrl == "" {
		apiUrl = viper.GetString("apiUrl")
	}

	apiKey := os.Getenv("RUNPOD_API_KEY")
	if apiKey == "" {
		apiKey = viper.GetString("apiKey")
	}

	// Check if the API key is present
	if apiKey == "" {
		fmt.Println("API key not found")
		return nil, errors.New("API key not found")
	}

	req, err := http.NewRequest("POST", apiUrl+"?api_key="+apiKey, bytes.NewBuffer(jsonValue))
	if err != nil {
		return
	}

	userAgent := "RunPod-CLI/" + Version + " (" + runtime.GOOS + "; " + runtime.GOARCH + ")"

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{Timeout: time.Second * 10}
	return client.Do(req)
}
