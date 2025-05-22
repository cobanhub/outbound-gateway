package outbound

import (
	"errors"
	"net/http"

	"github.com/cobanhub/madakaripura/internal/client"
	config "github.com/cobanhub/madakaripura/internal/integration_config"
	"github.com/cobanhub/madakaripura/internal/services/mapper"
)

func (o *Outbound) HandleJson(integrationName string, coreRequest map[string]interface{}, reqHeader http.Header) (map[string]interface{}, error) {
	// Validate the integration name

	cfg, err := config.GetIntegrationConfig(integrationName)
	if err != nil {
		// http.Error(w, "Integration not found", http.StatusBadRequest)
		return nil, errors.New("Integration not found")
	}

	thirdPartyRequest, thirdPartyHeadersReq, err := mapper.MapRequestWithHeaders(cfg, coreRequest, reqHeader)
	if err != nil {
		// http.Error(w, "Mapping request failed", http.StatusInternalServerError)
		return nil, errors.New("Mapping request failed")
	}

	thirdPartyResp, err := client.Send(thirdPartyRequest, thirdPartyHeadersReq, cfg)
	if err != nil {
		// http.Error(w, "Failed to contact third party", http.StatusBadGateway)
		return nil, errors.New("Failed to contact third party")
	}

	coreResponse, err := mapper.MapResponse(cfg, thirdPartyResp)
	if err != nil {
		// http.Error(w, "Mapping response failed", http.StatusInternalServerError)
		return nil, errors.New("Mapping response failed")
	}

	return coreResponse, nil
}
