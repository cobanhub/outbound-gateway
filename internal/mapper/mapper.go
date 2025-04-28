package mapper

import (
	"github.com/cobanhub/outbound-gateway/internal/client"
	"github.com/cobanhub/outbound-gateway/internal/config"
)

func MapRequest(cfg *config.IntegrationConfig, coreReq map[string]interface{}) (map[string]interface{}, error) {

	mapped := make(map[string]interface{})
	for k, v := range cfg.RequestMapping {
		mapped[k] = coreReq[v]
	}
	return mapped, nil
}

func MapResponse(cfg *config.IntegrationConfig, thirdResp client.GatewayResponse) (map[string]interface{}, error) {
	// Map success/fail first based on 3rd party response code
	mapped := map[string]interface{}{
		"status": "failed", // default
	}
	if code, ok := thirdResp.Body["code"].(string); ok {
		if result, exists := cfg.CodeMapping[code]; exists {
			mapped["status"] = result
		}
	}
	// Map other fields
	for k, v := range cfg.ResponseMapping {
		mapped[k] = thirdResp.Body[v]
	}
	return mapped, nil
}
