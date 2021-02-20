package hackernews

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableHackernewsJob(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hackernews_job",
		Description: "Latest 200 Job stories.",
		List: &plugin.ListConfig{
			Hydrate: hydrateList("jobstories"),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getItem,
		},
		Columns: itemCols(),
	}
}
