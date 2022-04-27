package hackernews

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableHackernewsBest(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hackernews_best",
		Description: "Best 500 stories.",
		List: &plugin.ListConfig{
			Hydrate: hydrateList("beststories"),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getItem,
		},
		Columns: itemCols(),
	}
}
