package mapper

import (
	"errors"
	"strings"

	"github.com/cobanhub/outbound-gateway/internal/config"
)

func MapRequest(coreRequest map[string]interface{}, integration *config.IntegrationConfig) (map[string]interface{}, error) {
	mapped := make(map[string]interface{})
	for targetField, sourceField := range integration.RequestMapping {
		value, found := getNestedField(coreRequest, sourceField)
		if !found {
			return nil, errors.New("missing field: " + sourceField)
		}
		mapped[targetField] = value
	}
	return mapped, nil
}

func MapResponse(thirdPartyResp map[string]interface{}, integration *config.IntegrationConfig) (map[string]interface{}, error) {
	mapped := make(map[string]interface{})
	for targetField, sourceField := range integration.ResponseMapping {
		value, found := getNestedField(thirdPartyResp, sourceField)
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
