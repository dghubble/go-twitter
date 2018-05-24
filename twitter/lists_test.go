package twitter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListsService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	expected := &[]List{
		{
			User:            User{ID: 123, Following: false},
			Slug:            "meetup-20100301",
			Name:            "meetup-20100301",
			CreatedAt:       "Sat Feb 27 21:39:24 +0000 2010",
			URI:             "/twitterapi/meetup-20100301",
			SubscriberCount: 147,
			IDStr:           "8044403",
			MemberCount:     116,
			Mode:            "public",
			ID:              8044403,
			FullName:        "@twitterapi/meetup-20100301",
			Description:     "Guests attending the Twitter meetup on 1 March 2010 at the @twoffice",
		},
		{
			User:            User{ID: 456, Following: true},
			Slug:            "team",
			Name:            "team",
			CreatedAt:       "Wed Nov 04 01:24:28 +0000 2009",
			URI:             "/twitterapi/team",
			SubscriberCount: 277,
			IDStr:           "2031945",
			MemberCount:     20,
			Mode:            "public",
			ID:              2031945,
			FullName:        "@twitterapi/team",
			Description:     "",
		},
	}

	query := map[string]string{"screen_name": "dghubble"}

	mux.HandleFunc("/1.1/lists/list.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, query, r)
		w.Header().Set("Content-Type", "application/json")
		expectedJSON, err := json.Marshal(expected)
		assert.Nil(t, err)
		fmt.Fprintf(w, string(expectedJSON))
	})

	client := NewClient(httpClient)

	params := &ListsListParams{
		ScreenName: "dghubble",
	}

	result, _, err := client.Lists.List(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestListsService_Members(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	expected := &Members{
		Users:             []User{{ID: 123, Following: false}, {ID: 456, Following: true}},
		NextCursor:        1516837838944119498,
		NextCursorStr:     "1516837838944119498",
		PreviousCursor:    -1516924983503961435,
		PreviousCursorStr: "-1516924983503961435",
	}

	query := map[string]string{
		"owner_screen_name": "dghubble",
		"slug":              "team",
	}

	mux.HandleFunc("/1.1/lists/members.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, query, r)
		w.Header().Set("Content-Type", "application/json")
		expectedJSON, err := json.Marshal(expected)
		assert.Nil(t, err)
		fmt.Fprintf(w, string(expectedJSON))
	})

	client := NewClient(httpClient)

	params := &ListsMembersParams{
		OwnerScreenName: "dghubble",
		Slug:            "team",
	}

	result, _, err := client.Lists.Members(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestListsService_MembersShow(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	expected := &User{
		ID:           123,
		IDStr:        "123",
		IsTranslator: false,
		Following:    true,
		ScreenName:   "froginthevalley",
	}

	query := map[string]string{
		"slug":              "team",
		"owner_screen_name": "dghubble",
		"screen_name":       "froginthevalley",
	}

	mux.HandleFunc("/1.1/lists/members/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, query, r)
		w.Header().Set("Content-Type", "application/json")
		expectedJSON, err := json.Marshal(expected)
		assert.Nil(t, err)
		fmt.Fprintf(w, string(expectedJSON))
	})

	client := NewClient(httpClient)

	params := &ListsMembersShowParams{
		Slug:            "team",
		OwnerScreenName: "dghubble",
		ScreenName:      "froginthevalley",
	}

	result, _, err := client.Lists.MembersShow(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestListsService_Memberships(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	expected := &Memberships{
		Lists: []List{
			{
				User:            User{ID: 123, Following: false},
				Slug:            "meetup-20100301",
				Name:            "meetup-20100301",
				CreatedAt:       "Sat Feb 27 21:39:24 +0000 2010",
				URI:             "/twitterapi/meetup-20100301",
				SubscriberCount: 147,
				IDStr:           "8044403",
				MemberCount:     116,
				Mode:            "public",
				ID:              8044403,
				FullName:        "@twitterapi/meetup-20100301",
				Description:     "Guests attending the Twitter meetup on 1 March 2010 at the @twoffice",
			},
			{
				User:            User{ID: 456, Following: true},
				Slug:            "team",
				Name:            "team",
				CreatedAt:       "Wed Nov 04 01:24:28 +0000 2009",
				URI:             "/twitterapi/team",
				SubscriberCount: 277,
				IDStr:           "2031945",
				MemberCount:     20,
				Mode:            "public",
				ID:              2031945,
				FullName:        "@twitterapi/team",
				Description:     "",
			},
		},
		NextCursor:        1516837838944119498,
		NextCursorStr:     "1516837838944119498",
		PreviousCursor:    -1516924983503961435,
		PreviousCursorStr: "-1516924983503961435",
	}

	query := map[string]string{
		"screen_name": "dghubble",
		"cursor":      "-1",
	}

	mux.HandleFunc("/1.1/lists/memberships.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, query, r)
		w.Header().Set("Content-Type", "application/json")
		expectedJSON, err := json.Marshal(expected)
		assert.Nil(t, err)
		fmt.Fprintf(w, string(expectedJSON))
	})

	client := NewClient(httpClient)

	params := &ListsMembershipsParams{
		ScreenName: "dghubble",
		Cursor:     -1,
	}

	result, _, err := client.Lists.Memberships(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
