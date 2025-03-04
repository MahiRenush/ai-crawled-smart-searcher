package main

import (
	"github.comcast.com/ciec-labweek/comcast-one/server/search"
)

func main() {
	search.OpenDB()
	search.WriteDB()
	// search.ReadDBStream()
	search.CloseDB()
}
