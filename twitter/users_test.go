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
		assertParams(t, map[string]string{"screen_name": "xkcdComic"}, r)
		fmt.Fprintf(w, `{"name": "XKCD Comic", "favourites_count": 2}`)
	})

	client := NewClient(httpClient)
	user, _, err := client.Users.Show(&UserShowParams{ScreenName: "xkcdComic"})
	if err != nil {
		t.Errorf("Users.Show unexpected error %v", err)
	}
	expected := &User{Name: "XKCD Comic", FavouritesCount: 2}
	if !reflect.DeepEqual(expected, user) {
		t.Errorf("Users.Show expected:\n%+v, got:\n %+v", expected, user)
	}
}

func TestUserService_Lookup(t *testing.T) {
	httpClient, mux, server := testServer() //ver(`[{"name": "Foo"}, {"name": "Bar"}]`)
	defer server.Close()

	mux.HandleFunc("/1.1/users/lookup.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertParamsWithDups(t, map[string][]string{"screen_name": []string{"foo", "bar"}}, r)
		fmt.Fprintf(w, `[{"name": "Foo"}, {"name": "Bar"}]`)
	})

	client := NewClient(httpClient)
	users, _, err := client.Users.Lookup(&UserLookupParams{ScreenName: []string{"foo", "bar"}})
	if err != nil {
		t.Errorf("Users.Lookup unexpected error %v", err)
	}
	expected := []User{User{Name: "Foo"}, User{Name: "Bar"}}
	if !reflect.DeepEqual(expected, users) {
		t.Errorf("Users.Lookup expected:\n%+v, got:\n %+v", expected, users)
	}
}

func TestSearch(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/users/search.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertParams(t, map[string]string{"count": "11", "q": "news"}, r)
		fmt.Fprintf(w, `[{"name": "BBC"}, {"name": "BBC Breaking News"}]`)
	})

	client := NewClient(httpClient)
	users, _, err := client.Users.Search("news", &UserSearchParams{Count: 11})
	if err != nil {
		t.Errorf("Users.Search unexpected error %v", err)
	}
	expected := []User{User{Name: "BBC"}, User{Name: "BBC Breaking News"}}
	if !reflect.DeepEqual(expected, users) {
		t.Errorf("Users.Show expected:\n%+v, got:\n %+v", expected, users)
	}
}
