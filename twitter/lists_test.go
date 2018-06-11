package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListsService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/list.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"screen_name": "twitterapi"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"slug": "meetup-20100301", "uri": "/twitterapi/meetup-20100301"}]`)
	})

	client := NewClient(httpClient)
	params := &ListsListParams{ScreenName: "twitterapi"}
	lists, _, err := client.Lists.List(params)
	expected := []List{List{Slug: "meetup-20100301", URI: "/twitterapi/meetup-20100301"}}
	assert.Nil(t, err)
	assert.Equal(t, expected, lists)
}

func TestListsService_Members(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/members.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"slug": "team", "owner_screen_name": "twitterapi", "cursor": "-1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"users":[{"id": 14895163}],"next_cursor":1516837838944119498,"next_cursor_str":"1516837838944119498","previous_cursor":-1516924983503961435,"previous_cursor_str":"-1516924983503961435"}`)
	})

	client := NewClient(httpClient)
	params := &ListsMembersParams{Slug: "team", OwnerScreenName: "twitterapi", Cursor: -1}
	members, _, err := client.Lists.Members(params)
	expected := &Members{
		Users:             []User{User{ID: 14895163}},
		NextCursor:        1516837838944119498,
		NextCursorStr:     "1516837838944119498",
		PreviousCursor:    -1516924983503961435,
		PreviousCursorStr: "-1516924983503961435",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, members)
}

func TestListsService_MembersShow(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/members/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"slug": "team", "owner_screen_name": "twitterapi", "screen_name": "froginthevalley"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id": 657693, "screen_name": "froginthevalley"}`)
	})

	client := NewClient(httpClient)
	params := &ListsMembersShowParams{Slug: "team", OwnerScreenName: "twitterapi", ScreenName: "froginthevalley"}
	user, _, err := client.Lists.MembersShow(params)
	expected := &User{
		ID:         657693,
		ScreenName: "froginthevalley",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}

func TestListsService_Memberships(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/memberships.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"screen_name": "twitter", "cursor": "-1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"lists": [{"slug": "digital-marketing", "name": "Digital Marketing"}], "next_cursor":1516837838944119498,"next_cursor_str":"1516837838944119498","previous_cursor":-1516924983503961435,"previous_cursor_str":"-1516924983503961435"}`)
	})

	client := NewClient(httpClient)
	params := &ListsMembershipsParams{ScreenName: "twitter", Cursor: -1}
	memberships, _, err := client.Lists.Memberships(params)
	expected := &Membership{
		Lists:             []List{List{Slug: "digital-marketing", Name: "Digital Marketing"}},
		NextCursor:        1516837838944119498,
		NextCursorStr:     "1516837838944119498",
		PreviousCursor:    -1516924983503961435,
		PreviousCursorStr: "-1516924983503961435",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, memberships)
}

func TestListsService_Ownerships(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/ownerships.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"screen_name": "twitter", "count": "2"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"lists": [{"mode": "public", "name": "Official Twitter accts"}], "next_cursor":1516837838944119498,"next_cursor_str":"1516837838944119498","previous_cursor":-1516924983503961435,"previous_cursor_str":"-1516924983503961435"}`)
	})

	client := NewClient(httpClient)
	params := &ListsOwnershipsParams{ScreenName: "twitter", Count: 2}
	ownerships, _, err := client.Lists.Ownerships(params)
	expected := &Ownership{
		Lists:             []List{List{Mode: "public", Name: "Official Twitter accts"}},
		NextCursor:        1516837838944119498,
		NextCursorStr:     "1516837838944119498",
		PreviousCursor:    -1516924983503961435,
		PreviousCursorStr: "-1516924983503961435",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, ownerships)
}

func TestListsService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"slug": "team", "owner_screen_name": "twitter"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"full_name": "@twitter/team", "following": false, "member_count": 643}`)
	})

	client := NewClient(httpClient)
	params := &ListsShowParams{Slug: "team", OwnerScreenName: "twitter"}
	list, _, err := client.Lists.Show(params)
	expected := &List{
		FullName:    "@twitter/team",
		Following:   false,
		MemberCount: 643,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}

func TestListsService_Statuses(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/statuses.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"slug": "teams", "owner_screen_name": "MLS", "count": "1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"user": {"screen_name": "torontofc"}, "id": 245160944223793152, "text": "Create your own TFC ESQ by Movado Watch: http://t.co/W2tON3OK in support of @TeamUpFdn #TorontoFC #MLS"}]`)
	})

	client := NewClient(httpClient)
	params := &ListsStatusesParams{Slug: "teams", OwnerScreenName: "MLS", Count: 1}
	tweet, _, err := client.Lists.Statuses(params)
	expected := []Tweet{
		Tweet{
			ID:   245160944223793152,
			Text: "Create your own TFC ESQ by Movado Watch: http://t.co/W2tON3OK in support of @TeamUpFdn #TorontoFC #MLS",
			User: &User{
				ScreenName: "torontofc",
			},
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, tweet)
}

func TestListsService_Subscribers(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/subscribers.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"slug": "team", "owner_screen_name": "twitter", "skip_status": "true"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"users": [{"name": "Almissen665"}], "next_cursor":1516837838944119498,"next_cursor_str":"1516837838944119498","previous_cursor":-1516924983503961435,"previous_cursor_str":"-1516924983503961435"}`)
	})

	client := NewClient(httpClient)
	params := &ListsSubscribersParams{Slug: "team", OwnerScreenName: "twitter", SkipStatus: Bool(true)}
	subscribers, _, err := client.Lists.Subscribers(params)
	expected := &Subscribers{
		Users:             []User{User{Name: "Almissen665"}},
		NextCursor:        1516837838944119498,
		NextCursorStr:     "1516837838944119498",
		PreviousCursor:    -1516924983503961435,
		PreviousCursorStr: "-1516924983503961435",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, subscribers)
}

func TestListsService_SubscribersShow(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/subscribers/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"slug": "team", "owner_screen_name": "twitter", "screen_name": "episod"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"name": "Taylor Singletary", "screen_name": "episod"}`)
	})

	client := NewClient(httpClient)
	params := &ListsSubscribersShowParams{Slug: "team", OwnerScreenName: "twitter", ScreenName: "episod"}
	user, _, err := client.Lists.SubscribersShow(params)
	expected := &User{
		Name:       "Taylor Singletary",
		ScreenName: "episod",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}

func TestListsService_Subscriptions(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/subscriptions.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"cursor": "-1", "screen_name": "episod", "count": "5"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"lists": [{"slug": "team", "name": "team", "uri": "/TwitterEng/team"}], "next_cursor":1516837838944119498,"next_cursor_str":"1516837838944119498","previous_cursor":-1516924983503961435,"previous_cursor_str":"-1516924983503961435"}`)
	})

	client := NewClient(httpClient)
	params := &ListsSubscriptionsParams{Cursor: -1, ScreenName: "episod", Count: 5}
	subscriptions, _, err := client.Lists.Subscriptions(params)
	expected := &Subscribed{
		Lists:             []List{List{Slug: "team", Name: "team", URI: "/TwitterEng/team"}},
		NextCursor:        1516837838944119498,
		NextCursorStr:     "1516837838944119498",
		PreviousCursor:    -1516924983503961435,
		PreviousCursorStr: "-1516924983503961435",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, subscriptions)
}

func TestListsService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/create.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"name": "Goonies", "mode": "public", "description": "For life"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"slug": "goonies", "name": "Goonies", "description": "For life"}`)
	})

	client := NewClient(httpClient)
	params := &ListsCreateParams{Mode: "public", Description: "For life"}
	list, _, err := client.Lists.Create("Goonies", params)
	expected := &List{
		Slug:        "goonies",
		Name:        "Goonies",
		Description: "For life",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}

func TestListsService_Destroy(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"owner_screen_name": "kurrik", "slug": "goonies"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"slug": "goonies", "name": "Goonies", "full_name": "@kurrik/goonies"}`)
	})

	client := NewClient(httpClient)
	params := &ListsDestroyParams{OwnerScreenName: "kurrik", Slug: "goonies"}
	list, _, err := client.Lists.Destroy(params)
	expected := &List{
		Slug:     "goonies",
		Name:     "Goonies",
		FullName: "@kurrik/goonies",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}

func TestListsService_MembersCreate(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/members/create.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"slug": "team", "owner_screen_name": "twitter", "screen_name": "kurrik"}, r)
		w.Header().Set("Content-Type", "application/json")
	})

	client := NewClient(httpClient)
	params := &ListsMembersCreateParams{Slug: "team", OwnerScreenName: "twitter", ScreenName: "kurrik"}
	_, err := client.Lists.MembersCreate(params)
	assert.Nil(t, err)
}

func TestListsService_MembersCreateAll(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/members/create_all.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"list_id": "23", "screen_name": "rsarver,episod,jasoncosta"}, r)
		w.Header().Set("Content-Type", "application/json")
	})

	client := NewClient(httpClient)
	params := &ListsMembersCreateAllParams{ListID: 23, ScreenName: "rsarver,episod,jasoncosta"}
	_, err := client.Lists.MembersCreateAll(params)
	assert.Nil(t, err)
}

func TestListsService_MembersDestroy(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/members/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"screen_name": "episod", "slug": "cool_people", "owner_screen_name": "twitter"}, r)
		w.Header().Set("Content-Type", "application/json")
	})

	client := NewClient(httpClient)
	params := &ListsMembersDestroyParams{ScreenName: "episod", Slug: "cool_people", OwnerScreenName: "twitter"}
	_, err := client.Lists.MembersDestroy(params)
	assert.Nil(t, err)
}

func TestListsService_DestroyAll(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/members/destroy_all.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"screen_name": "rsarver,episod,jasoncosta,theseancook,kurrik,froginthevalley", "list_id": "23"}, r)
		w.Header().Set("Content-Type", "application/json")
	})

	client := NewClient(httpClient)
	params := &ListsMembersDestroyAllParams{ScreenName: "rsarver,episod,jasoncosta,theseancook,kurrik,froginthevalley", ListID: 23}
	_, err := client.Lists.MembersDestroyAll(params)
	assert.Nil(t, err)
}

func TestListsService_SubscribersCreate(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/subscribers/create.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"slug": "team", "owner_screen_name": "twitter"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"following": false, "id_str": "574"}`)
	})

	client := NewClient(httpClient)
	params := &ListsSubscribersCreateParams{Slug: "team", OwnerScreenName: "twitter"}
	list, _, err := client.Lists.SubscribersCreate(params)
	expected := &List{
		Following: false,
		IDStr:     "574",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}

func TestListsService_SubscribersDestroy(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/subscribers/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"slug": "team", "owner_screen_name": "twitterapi"}, r)
		w.Header().Set("Content-Type", "application/json")
	})

	client := NewClient(httpClient)
	params := &ListsSubscribersDestroyParams{Slug: "team", OwnerScreenName: "twitterapi"}
	_, err := client.Lists.SubscribersDestroy(params)
	assert.Nil(t, err)
}

func TestListsService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/lists/update.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"list_id": "1234", "mode": "public", "name": "Party Time"}, r)
		w.Header().Set("Content-Type", "application/json")
	})

	client := NewClient(httpClient)
	params := &ListsUpdateParams{ListID: 1234, Mode: "public", Name: "Party Time"}
	_, err := client.Lists.Update(params)
	assert.Nil(t, err)
}
