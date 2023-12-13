package hackernews

import (
	"context"

	"github.com/PaulRosset/go-hacknews"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func itemCols() []*plugin.Column {
	return []*plugin.Column{
		// Top columns
		{Name: "id", Type: proto.ColumnType_INT, Description: "The item's unique id."},
		// TODO - How can I set the default hydrate method?
		{Name: "title", Type: proto.ColumnType_STRING, Hydrate: getItem, Description: "The title of the story, poll or job. HTML."},
		{Name: "time", Type: proto.ColumnType_TIMESTAMP, Hydrate: getItem, Transform: transform.FromField("Time").Transform(transform.UnixToTimestamp), Description: "Timestamp when the item was created."},
		{Name: "by", Type: proto.ColumnType_STRING, Hydrate: getItem, Description: "The username of the item's author."},
		{Name: "score", Type: proto.ColumnType_INT, Hydrate: getItem, Description: "The story's score, or the votes for a pollopt."},
		// Other columns
		{Name: "dead", Type: proto.ColumnType_BOOL, Hydrate: getItem, Description: "True if the item is dead."},
		{Name: "deleted", Type: proto.ColumnType_BOOL, Hydrate: getItem, Description: "True if the item is deleted."},
		{Name: "descendants", Type: proto.ColumnType_INT, Hydrate: getItem, Description: "In the case of stories or polls, the total comment count."},
		{Name: "kids", Type: proto.ColumnType_JSON, Hydrate: getItem, Description: "The ids of the item's comments, in ranked display order."},
		{Name: "parent", Type: proto.ColumnType_INT, Hydrate: getItem, Description: "The comment's parent: either another comment or the relevant story."},
		{Name: "parts", Type: proto.ColumnType_JSON, Hydrate: getItem, Description: "A list of related pollopts, in display order."},
		{Name: "poll", Type: proto.ColumnType_INT, Hydrate: getItem, Description: "The pollopt's associated poll."},
		{Name: "text", Type: proto.ColumnType_STRING, Hydrate: getItem, Description: "The comment, story or poll text. HTML."},
		{Name: "type", Type: proto.ColumnType_STRING, Hydrate: getItem, Description: "The type of item. One of \"job\", \"story\", \"comment\", \"poll\", or \"pollopt\"."},
		{Name: "url", Type: proto.ColumnType_STRING, Hydrate: getItem, Description: "The URL of the story."},
	}
}

func hydrateList(listType string) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
		init := hacknews.Initializer{Story: listType, NbPosts: 0}
		codes, err := init.GetCodesStory()
		if err != nil {
			plugin.Logger(ctx).Error("hackernews_*.list", "query_error", err)
			return nil, err
		}
		for _, id := range codes {
			d.StreamListItem(ctx, &hacknews.Post{Id: id})
		}
		return nil, nil
	}
}

func getItem(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var id int
	if h.Item != nil {
		post := h.Item.(*hacknews.Post)
		id = post.Id
	} else {
		quals := d.EqualsQuals
		id = int(quals["id"].GetInt64Value())
	}
	init := hacknews.Initializer{Story: "_ignored_", NbPosts: 1}
	posts, err := init.GetPostStory([]int{int(id)})
	if err != nil {
		plugin.Logger(ctx).Error("hackernews_top.getTopStory", "query_error", err)
		return nil, err
	}
	if len(posts) <= 0 {
		return hacknews.Post{Id: id}, nil
	}
	return posts[0], nil
}
