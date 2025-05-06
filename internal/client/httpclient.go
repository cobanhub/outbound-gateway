package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
func Send(payload map[string]interface{}, headers map[string]string, integration *config.IntegrationConfig) (GatewayResponse, error) {
	bodyBytes, err := json.Marshal(payload)

	if err != nil {
		return GatewayResponse{}, err
	}

	result, err := cb.Execute(func() (interface{}, error) {
		req, err := http.NewRequest(integration.Method, integration.Endpoint, bytes.NewReader(bodyBytes))
		if err != nil {
			return nil, err
		}

		if integration != nil {
			if headers != nil {
				for key, value := range headers {
					req.Header.Set(key, value)
				}
			}
			// if integration.QueryParams != nil {
			// 	q := req.URL.Query()
			// 	for key, value := range integration.QueryParams {
			// 		q.Add(key, value)
			// 	}
			// 	req.URL.RawQuery = q.Encode()
			// }
		}

		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		var resp *http.Response
		for attempts := 0; attempts < integration.RetryCount; attempts++ {
			resp, err = client.Do(req)
			if err == nil {
				break
			}
			time.Sleep(time.Duration(integration.RetryInterval) * time.Second)
		}
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var respData map[string]interface{}
		if err := json.Unmarshal(respBytes, &respData); err != nil {
			return nil, fmt.Errorf("failed to parse response: %w", err)
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
