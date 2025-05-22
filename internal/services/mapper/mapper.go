package mapper

import (
	"errors"
	"net/http"
	"strings"

	"github.com/cobanhub/madakaripura/internal/client"
	config "github.com/cobanhub/madakaripura/internal/integration_config"
)

func MapRequestWithHeaders(cfg *config.IntegrationConfig, coreReq map[string]interface{}, headersReq http.Header) (req map[string]interface{}, header map[string]string, err error) {
	req = make(map[string]interface{})
	header = make(map[string]string)

	// Map the request body
	for targetField, sourceField := range cfg.RequestMapping {
		value, found := getNestedField(coreReq, sourceField)
		if !found {
			return nil, nil, errors.New("missing field: " + sourceField)
		}
		req[targetField] = value
	}

	// Map the headers
	for targetField, sourceField := range cfg.HeadersMapping {
		value := headersReq.Get(sourceField)
		if value == "" {
			return nil, nil, errors.New("missing field: " + sourceField)
		}
		header[targetField] = value
	}

	return req, header, nil
}

func MapRequest(cfg *config.IntegrationConfig, coreReq map[string]interface{}) (map[string]interface{}, error) {
	mapped := make(map[string]interface{})
	for targetField, sourceField := range cfg.RequestMapping {
		value, found := getNestedField(coreReq, sourceField)
		if !found {
			return nil, errors.New("missing field: " + sourceField)
		}
		mapped[targetField] = value
	}
	return mapped, nil
}

func MapResponse(integration *config.IntegrationConfig, thirdPartyResp client.GatewayResponse) (map[string]interface{}, error) {
	mapped := make(map[string]interface{})
	for targetField, sourceField := range integration.ResponseMapping {
		value, found := getNestedField(thirdPartyResp.Body, sourceField)
		if !found {
			return nil, errors.New("missing field: " + sourceField)
		}
		mapped[targetField] = value
	}
	return mapped, nil
}
func getNestedField(data map[string]interface{}, path string) (interface{}, bool) {
	keys := splitPath(path)
	var current interface{} = data
	for _, k := range keys {
		if m, ok := current.(map[string]interface{}); ok {
			current, ok = m[k]
			if !ok {
				return nil, false
			}
		} else {
			return nil, false
		}
	}
	return current, true
}

func splitPath(path string) []string {
	return strings.Split(path, ".")
}
