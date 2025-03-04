package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.comcast.com/ciec-labweek/comcast-one/server/search"
)

var empty []search.Event = []search.Event{}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["q"]
	if !ok || len(params[0]) < 1 {
		result := search.ReadDBStream()
		bytes, mErr := json.Marshal(result)
		if nil != mErr {
			fmt.Println("Marshal error")
		}
		w.Write(bytes)
		return
	}

	result, err := search.SearchForQuery(params[0])
	if err != nil {
		log.Println("Error while Search", err)
	}
	if result == nil {
		emptyData, mErr := json.Marshal(empty)
		if nil != mErr {
			fmt.Println("Error marshalling")
		}
		w.Write([]byte(emptyData))
	}
	defer search.CloseDB()
	w.Write(result)
}

func SearchBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	result := search.ReadBookmarks()
	bytes, mErr := json.Marshal(result)
	if nil != mErr {
		fmt.Println("Marshal error")
	}
	w.Write(bytes)
	return
}

func CreateBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bookmark search.Bookmark
	json.NewDecoder(r.Body).Decode(&bookmark)
	search.AddBookmarks(bookmark.Title, bookmark.Url)
	json.NewEncoder(w).Encode(bookmark)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	body, _ := ioutil.ReadAll(r.Body)
	w.Write([]byte(body))
}
