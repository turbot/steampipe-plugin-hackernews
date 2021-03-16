package hackernews

import (
	"context"

	"github.com/PaulRosset/go-hacknews"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableHackernewsItem(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hackernews_item",
		Description: "Stories, comments, jobs, Ask HNs and even polls are just items. This table includes the most recent items posted to Hacker News.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    itemList,
		},
		Columns: itemCols(),
	}
}

func itemList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	item, err := getItem(ctx, d, h)
	if err != nil {
		return nil, err
	}
	postItem := item.(hacknews.Post)
	d.StreamListItem(ctx, &postItem)

	return nil, nil
}
