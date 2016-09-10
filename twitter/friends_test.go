package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFriendService_Ids(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friends/ids.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"user_id": "623265148", "count": "5", "cursor": "1541399463850369479"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"ids":[3286770860,2560569758,403722417,754745643399282688,2498920249],"next_cursor":1540043710791833611,"next_cursor_str":"1540043710791833611","previous_cursor":-1541130110610547546,"previous_cursor_str":"-1541130110610547546"}`)
	})
	expected := &FriendIDs{
		IDs:               []int64{3286770860, 2560569758, 403722417, 754745643399282688, 2498920249},
		NextCursor:        1540043710791833611,
		NextCursorStr:     "1540043710791833611",
		PreviousCursor:    -1541130110610547546,
		PreviousCursorStr: "-1541130110610547546",
	}

	client := NewClient(httpClient)
	params := &FriendIDParams{
		UserID: 623265148,
		Count:  5,
		Cursor: 1541399463850369479,
	}
	friendIDs, _, err := client.Friends.IDs(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, friendIDs)
}

func TestFriendService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friends/list.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"screen_name": "dghubble", "count": "5", "cursor": "1516933260114270762", "skip_status": "true", "include_user_entities": "false"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"users": [{"id": 3286770860}], "next_cursor":1540043710791833611,"next_cursor_str":"1540043710791833611","previous_cursor":-1541130110610547546,"previous_cursor_str":"-1541130110610547546"}`)
	})
	expected := &Friends{
		Users:             []User{User{ID: 3286770860}},
		NextCursor:        1540043710791833611,
		NextCursorStr:     "1540043710791833611",
		PreviousCursor:    -1541130110610547546,
		PreviousCursorStr: "-1541130110610547546",
	}

	client := NewClient(httpClient)
	params := &FriendListParams{
		ScreenName:          "dghubble",
		Count:               5,
		Cursor:              1516933260114270762,
		SkipStatus:          Bool(true),
		IncludeUserEntities: Bool(false),
	}
	friends, _, err := client.Friends.List(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, friends)
}
