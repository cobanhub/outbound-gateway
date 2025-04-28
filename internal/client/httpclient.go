package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cobanhub/outbound-gateway/internal/config"

	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func init() {
	settings := gobreaker.Settings{
		Name:        "ThirdPartyAPICircuitBreaker",
		MaxRequests: 5,                // allowed during half-open
		Interval:    60 * time.Second, // refresh counts
		Timeout:     30 * time.Second, // how long to stay open
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRate := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 10 && failureRate >= 0.5
		},
	}
	cb = gobreaker.NewCircuitBreaker(settings)
}
func Send(payload map[string]interface{}, integration *config.IntegrationConfig) (GatewayResponse, error) {
	bodyBytes, _ := json.Marshal(payload)

	result, err := cb.Execute(func() (interface{}, error) {
		req, err := http.NewRequest(integration.Method, integration.Endpoint, bytes.NewReader(bodyBytes))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")

		if integration != nil && integration.Auth != nil {
			if strings.EqualFold(integration.Auth.Type, "Bearer") {
				req.Header.Set("Authorization", "Bearer "+integration.Auth.Token)
			}
		}

		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		var resp *http.Response
		maxRetries := 3
		for attempts := 0; attempts < maxRetries; attempts++ {
			resp, err = client.Do(req)
			if err == nil {
				break
			}
			time.Sleep(time.Duration(attempts+1) * time.Second)
		}
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var respData map[string]interface{}
		if err := json.Unmarshal(respBytes, &respData); err != nil {
			return nil, errors.New("Invalid 3rd party response")
		}

		return GatewayResponse{
			Body:   respData,
			Status: resp.StatusCode,
		}, nil
	})

	if err != nil {
		return GatewayResponse{}, err
	}

	return result.(GatewayResponse), nil
}
