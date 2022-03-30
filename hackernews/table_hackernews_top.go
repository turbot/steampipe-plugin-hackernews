package hackernews

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func tableHackernewsTop(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hackernews_top_story",
		Description: "Top 500 stories.",
		List: &plugin.ListConfig{
			Hydrate: hydrateList("topstories"),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getItem,
		},
		Columns: itemCols(),
	}
}
