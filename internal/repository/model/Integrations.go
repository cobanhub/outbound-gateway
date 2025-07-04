package model

type (
	ConfigurationModel struct {
		StoreConfigFilePath string
		Type                StorageTypeModel
	}

	StorageConfigModel struct {
	}

	StorageTypeModel int
	Integrations     struct {
		Integrations IntegrationConfigModel `yaml:"integrations"`
	}

	IntegrationConfigModel struct {
		Name               string            `yaml:"name"`
		MatchPath          string            `yaml:"matchPath"`
		Endpoint           string            `yaml:"endpoint"`
		Method             string            `yaml:"method"`
		RequestMapping     map[string]string `yaml:"requestMapping"`
		ResponseMapping    map[string]string `yaml:"responseMapping"`
		Auth               *AuthConfigModel  `yaml:"auth"`
		CodeMapping        map[string]string // third party codes => success/failed
		AdapterScript      string            // Lua adapter optional
		RetryCount         int               `yaml:"retryCount"`
		RetryInterval      int               `yaml:"retryInterval"`
		Timeout            int               `yaml:"timeout"`
		HeadersMapping     map[string]string `yaml:"headersMapping"`
		QueryParamsMapping map[string]string `yaml:"queryParamsMapping"`
	}

	AuthConfigModel struct {
		Type  string `yaml:"type"`
		Token string `yaml:"token"`
	}
)
