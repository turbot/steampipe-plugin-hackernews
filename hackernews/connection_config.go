package hackernews

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type hackernewsConfig struct {
	MaxItems *int `hcl:"max_items"`
}

func ConfigInstance() interface{} {
	return &hackernewsConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) hackernewsConfig {
	if connection == nil || connection.Config == nil {
		return hackernewsConfig{}
	}
	config, _ := connection.Config.(hackernewsConfig)
	return config
}
