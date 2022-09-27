package hackernews

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-hackernews",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		DefaultConcurrency: &plugin.DefaultConcurrencyConfig{
			TotalMaxConcurrency:   500,
			DefaultMaxConcurrency: 200,
		},
		TableMap: map[string]*plugin.Table{
			"hackernews_ask_hn":  tableHackernewsAskHn(ctx),
			"hackernews_best":    tableHackernewsBest(ctx),
			"hackernews_item":    tableHackernewsItem(ctx),
			"hackernews_job":     tableHackernewsJob(ctx),
			"hackernews_new":     tableHackernewsNew(ctx),
			"hackernews_show_hn": tableHackernewsShowHn(ctx),
			"hackernews_top":     tableHackernewsTop(ctx),
			"hackernews_user":    tableHackernewsUser(ctx),
		},
	}
	return p
}
