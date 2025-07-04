package controller

import (
	"github.com/cobanhub/lib/response"
	"github.com/cobanhub/lib/router"
)

func (a *API) HandleOutbound(ctx *router.Ctx) *response.JSONResponse {
	integrationName := ctx.GetParams("integration")
	resp, err := a.interactor.HandleJson(integrationName, ctx)
	if err != nil {
		return response.NewJSONResponse().SetCode(500).SetData(err)
	}

	return response.NewJSONResponse().SetData(resp).SetCode(200)
}
func (a *API) UploadConfigHandler(ctx *router.Ctx) *response.JSONResponse {
	// Retrieve the YAML file from the request
	fileHeader, err := ctx.FormFile("config_file")
	if err != nil {
		return response.NewJSONResponse().SetCode(400).SetData("Failed to retrieve file")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return response.NewJSONResponse().SetCode(400).SetData("Failed to open file")
	}

	resp, err := a.interactor.UploadConfigHandler(file)
	if err != nil {
		return response.NewJSONResponse().SetCode(500).SetData(err)
	}
	return response.NewJSONResponse().SetData(resp).SetCode(200)
}
