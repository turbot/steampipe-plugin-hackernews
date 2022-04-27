package hackernews

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableHackernewsShowHn(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hackernews_show_hn",
		Description: "Latest 200 Show HN stories.",
		List: &plugin.ListConfig{
			Hydrate: hydrateList("showstories"),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getItem,
		},
		Columns: itemCols(),
	}
}
