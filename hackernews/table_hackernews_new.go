package hackernews

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableHackernewsNew(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hackernews_new",
		Description: "Newest 500 stories.",
		List: &plugin.ListConfig{
			Hydrate: hydrateList("newstories"),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getItem,
		},
		Columns: itemCols(),
	}
}
