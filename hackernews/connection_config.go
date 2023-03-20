package hackernews

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type hackernewsConfig struct {
	MaxItems *int `cty:"max_items"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"max_items": {
		Type: schema.TypeInt,
	},
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
