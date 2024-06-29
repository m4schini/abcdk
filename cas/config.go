package cas

const (
	DefaultBucketName     = "objects"
	DefaultDatabaseName   = "cas"
	DefaultCollectionName = "objects"
)

type ClientConfig struct {
	BucketName     string `json:"bucketName"`
	DatabaseName   string `json:"databaseName"`
	CollectionName string `json:"collectionName"`
}

func applyClientConfig(config ClientConfig) ClientConfig {
	if config.BucketName == "" {
		config.BucketName = DefaultBucketName
	}
	if config.DatabaseName == "" {
		config.DatabaseName = DefaultDatabaseName
	}
	if config.CollectionName == "" {
		config.CollectionName = DefaultCollectionName
	}
	return config
}
