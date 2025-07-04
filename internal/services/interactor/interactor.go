package interactor

import (
	"mime/multipart"

	"github.com/cobanhub/lib/response"
	"github.com/cobanhub/lib/router"
)

type Interactor interface {
	HandleJson(integrationName string, ctx *router.Ctx) (map[string]interface{}, error)
	//TODO: add more handlers like HandleXML, HandleGrpc, HandleGraphQL etc.
	// HandleXML(integrationName string, coreRequest map[string]interface{}) (map[string]interface{}, error)
	// HandleCSV(integrationName string, coreRequest map[string]interface{}) (map[string]interface{}, error)
	// HandleFile(integrationName string, coreRequest map[string]interface{}) (map[string]interface{}, error)
	UploadConfigHandler(file multipart.File) (response.JSONResponse, error)
}
