package hackernews

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type hackernewsConfig struct {
	MaxItems *int `hcl:"max_items"`
}

func ConfigInstance() interface{} {
	return &hackernewsConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) hackernewsConfig {
	if connection == nil || connection.GetConfig() == nil {
		return hackernewsConfig{}
	}
	config, _ := connection.GetConfig().(hackernewsConfig)
	return config
}
