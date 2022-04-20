package hackernews

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

type User struct {
	ID        string `json:"id"`
	Created   int    `json:"created"`
	Karma     int    `json:"karma"`
	About     string `json:"about,omitempty"`
	Submitted []int  `json:"submitted"`
}

func tableHackernewsUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hackernews_user",
		Description: "Information about Hacker News registered users who have public activity.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    listUser,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The user's unique username. Case-sensitive. Required."},
			{Name: "created", Type: proto.ColumnType_STRING, Description: "Creation timestamp of the user."},
			{Name: "karma", Type: proto.ColumnType_INT, Description: "The user's karma."},
			{Name: "about", Type: proto.ColumnType_STRING, Description: "The user's optional self-description. HTML."},
			{Name: "submitted", Type: proto.ColumnType_JSON, Description: "List of the user's stories, polls and comments."},
		},
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var item User
	var id string

	if h.Item != nil {
		user := h.Item.(*User)
		id = user.ID
	} else {
		quals := d.KeyColumnQuals
		id = quals["id"].GetStringValue()
	}

	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/user/" + id + ".json")
	if err != nil {
		plugin.Logger(ctx).Error("hackernews_user.getUser", "query_error", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		plugin.Logger(ctx).Error("hackernews_user.getUser", "read_body_error", err)
		return nil, err
	}

	// TODO - Should this return a not found error?
	if string(body) == "null" {
		return nil, nil
	}

	err = json.Unmarshal(body, &item)
	if err != nil {
		plugin.Logger(ctx).Error("hackernews_user.getUser", "unmarshal_json_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, item)

	return nil, nil
}
