package interactor

import "net/http"

type Interactor interface {
	HandleJson(integrationName string, coreRequest map[string]interface{}, reqHeader http.Header) (map[string]interface{}, error)
	//TODO: add more handlers like HandleXML, HandleGrpc, HandleGraphQL etc.
	// HandleXML(integrationName string, coreRequest map[string]interface{}) (map[string]interface{}, error)
	// HandleCSV(integrationName string, coreRequest map[string]interface{}) (map[string]interface{}, error)
	// HandleFile(integrationName string, coreRequest map[string]interface{}) (map[string]interface{}, error)
	UploadConfigHandler(fileContent []byte) error
}
