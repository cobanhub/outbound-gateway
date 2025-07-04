package request

type (
	ConfigurationRequest struct {
		StoreConfigFilePath string
		Type                StorageTypeRequest
	}

	StorageConfigRequest struct {
	}

	StorageTypeRequest int
	Integrations       struct {
		Integrations IntegrationConfigRequest `yaml:"integrations"`
	}

	IntegrationConfigRequest struct {
		Name               string             `yaml:"name"`
		MatchPath          string             `yaml:"matchPath"`
		Endpoint           string             `yaml:"endpoint"`
		Method             string             `yaml:"method"`
		RequestMapping     map[string]string  `yaml:"requestMapping"`
		ResponseMapping    map[string]string  `yaml:"responseMapping"`
		Auth               *AuthConfigRequest `yaml:"auth"`
		CodeMapping        map[string]string  // third party codes => success/failed
		AdapterScript      string             // Lua adapter optional
		RetryCount         int                `yaml:"retryCount"`
		RetryInterval      int                `yaml:"retryInterval"`
		Timeout            int                `yaml:"timeout"`
		HeadersMapping     map[string]string  `yaml:"headersMapping"`
		QueryParamsMapping map[string]string  `yaml:"queryParamsMapping"`
	}

	AuthConfigRequest struct {
		Type  string `yaml:"type"`
		Token string `yaml:"token"`
	}
)
