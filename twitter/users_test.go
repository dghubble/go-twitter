package twitter

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUserService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/users/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"screen_name": "xkcdComic"}, r)
		fmt.Fprintf(w, `{"name": "XKCD Comic", "favourites_count": 2}`)
	})

	client := NewClient(httpClient)
	user, _, err := client.Users.Show(&UserShowParams{ScreenName: "xkcdComic"})
	if err != nil {
		t.Errorf("Users.Show error %+v", err)
	}
	expected := &User{Name: "XKCD Comic", FavouritesCount: 2}
	if !reflect.DeepEqual(expected, user) {
		t.Errorf("Users.Show expected:\n%+v, got:\n %+v", expected, user)
	}
}

func TestUserService_LookupWithIds(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/users/lookup.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"user_id": "113419064,623265148"}, r)
		fmt.Fprintf(w, `[{"screen_name": "golang"}, {"screen_name": "dghubble"}]`)
	})

	client := NewClient(httpClient)
	users, _, err := client.Users.Lookup(&UserLookupParams{UserID: []int64{113419064, 623265148}})
	if err != nil {
		t.Errorf("Users.Lookup error %v", err)
	}
	expected := []User{User{ScreenName: "golang"}, User{ScreenName: "dghubble"}}
	if !reflect.DeepEqual(expected, users) {
		t.Errorf("Users.Lookup expected:\n%+v, got:\n %+v", expected, users)
	}
}

func TestUserService_LookupWithScreenNames(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/users/lookup.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"screen_name": "foo,bar"}, r)
		fmt.Fprintf(w, `[{"name": "Foo"}, {"name": "Bar"}]`)
	})

	client := NewClient(httpClient)
	users, _, err := client.Users.Lookup(&UserLookupParams{ScreenName: []string{"foo", "bar"}})
	if err != nil {
		t.Errorf("Users.Lookup error %v", err)
	}
	expected := []User{User{Name: "Foo"}, User{Name: "Bar"}}
	if !reflect.DeepEqual(expected, users) {
		t.Errorf("Users.Lookup expected:\n%+v, got:\n %+v", expected, users)
	}
}

func TestUserService_Search(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/users/search.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"count": "11", "q": "news"}, r)
		fmt.Fprintf(w, `[{"name": "BBC"}, {"name": "BBC Breaking News"}]`)
	})

	client := NewClient(httpClient)
	users, _, err := client.Users.Search("news", &UserSearchParams{Query: "override me", Count: 11})
	if err != nil {
		t.Errorf("Users.Search error %v", err)
	}
	expected := []User{User{Name: "BBC"}, User{Name: "BBC Breaking News"}}
	if !reflect.DeepEqual(expected, users) {
		t.Errorf("Users.Search expected:\n%+v, got:\n %+v", expected, users)
	}
}

func TestUserService_SearchHandlesNilParams(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/users/search.json", func(w http.ResponseWriter, r *http.Request) {
		assertQuery(t, map[string]string{"q": "news"}, r)
	})
	client := NewClient(httpClient)
	client.Users.Search("news", nil)
}
