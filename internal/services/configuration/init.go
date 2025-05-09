package configuration

type (
	// Configuration struct holds the configuration for the integration.
	// Type of the configuration storage (e.g., local, database, cloud storage, shared drive)
	// This can be used to determine where to store the configuration files
	// and how to access them.
	// For example, if you want to store the configuration in a database,
	// you might want to add a field for the database connection string.
	Configuration struct {
		StoreConfigFilePath string
		Type                StorageType
	}

	StorageConfig struct {
	}

	StorageType int
)

const (
	Local StorageType = iota
	Database
	Cloud
	SharedDrive
)

// func (s StorageType) String() string {
// 	switch s {
// 	case Local:
// 		return "local"
// 	case Database:
// 		return "database"
// 	case Cloud:
// 		return "cloud"
// 	case SharedDrive:
// 		return "shared_drive"
// 	default:
// 		return "unknown"
// 	}
// }
