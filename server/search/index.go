package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/blevesearch/bleve/v2"
)

const testIdx = "bleve.index"

func createIndex(path string) (bleve.Index, error) {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(path, mapping)
	if err != nil {
		fmt.Println("INDEX new path mapping ERROR:", err)
		index, anerr := bleve.Open(path)
		if anerr != nil {
			fmt.Println("INDEX open ERROR:", anerr)
			return nil, anerr
		}
		return index, nil
	}
	return index, nil
}

func (e *Event) index(index bleve.Index) error {
	err := index.Index(strconv.Itoa(e.ID), e)
	return err
}
func indexEvents(idx bleve.Index, eventList []Event) {
	for _, event := range eventList {
		event.index(idx)
	}
}

func SearchForQuery(searchTerm string) ([]byte, error) {
	var searchedEvents []Event
	eventList := ReadDBStream()
	fmt.Println("eventList", eventList)
	idx, err := createIndex(testIdx)
	if err != nil {
		fmt.Println("INDEX create error", err)
	}
	indexEvents(idx, eventList)
	query := bleve.NewMatchQuery(searchTerm)
	searchRequest := bleve.NewSearchRequest(query)
	searchResults, err := idx.Search(searchRequest)
	if err != nil {
		fmt.Println("Search ERROR:", err)
		return nil, err
	}
	differr := idx.Close()
	if differr != nil {
		fmt.Println(err)
	}
	if searchResults.Status.Successful == 1 && searchResults.Total > 0 {
		for i := 0; i < int(searchResults.Total); i++ {
			for _, event := range eventList {
				id, _ := strconv.Atoi(searchResults.Hits[i].ID)
				if id == event.ID {
					searchedEvents = append(searchedEvents, event)
				}
			}
		}
		fmt.Println("Searched Results Total -->", searchResults.Total)
		fmt.Println("Searched Events --> ", searchedEvents)
		return convStructToByte(searchedEvents), nil
	}
	defer idxDestroy()
	return nil, errors.New("No results obtained")
}

func convStructToByte(events []Event) (out []byte) {
	out, err := json.Marshal(events)
	if err != nil {
		fmt.Println("Marshall ERROR: ", err)
	}
	return out
}

func idxDestroy() {
	os.RemoveAll(testIdx)
}
