package hackernews

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PaulRosset/go-hacknews"

	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func tableHackernewsItem(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hackernews_item",
		Description: "Stories, comments, jobs, Ask HNs and even polls are just items. This table includes the most recent items posted to Hacker News.",
		List: &plugin.ListConfig{
			Hydrate: itemList,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getItem,
		},
		Columns: itemCols(),
	}
}

func itemList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/maxitem.json")
	if err != nil {
		plugin.Logger(ctx).Error("hackernews_item.itemList", "query_error", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		plugin.Logger(ctx).Error("hackernews_item.itemList", "read_body_error", err)
		return nil, err
	}

	max64, err := strconv.ParseInt(string(body), 10, 0)
	if err != nil {
		plugin.Logger(ctx).Error("hackernews_item.itemList", "parse_max_item_error", err)
		return nil, err
	}
	max := int(max64)

	// default to 5000 max items, but settable in config
	limit := 5000
	config := GetConfig(d.Connection)
	
	if config.MaxItems != nil {
		limit = *config.MaxItems
	}

	for i := max; i >= max-limit; i-- {
		d.StreamListItem(ctx, &hacknews.Post{Id: i})
	}

	return nil, nil
}
