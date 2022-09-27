package main

import (
	"github.com/turbot/steampipe-plugin-hackernews/hackernews"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: hackernews.Plugin})
}
