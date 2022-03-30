package hackernews

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func tableHackernewsAskHn(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hackernews_ask_hn",
		Description: "Latest 200 Ask HN stories.",
		List: &plugin.ListConfig{
			Hydrate: hydrateList("askstories"),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getItem,
		},
		Columns: itemCols(),
	}
}
