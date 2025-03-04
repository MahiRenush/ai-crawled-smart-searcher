package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.comcast.com/ciec-labweek/comcast-one/server/handler"
	"github.comcast.com/ciec-labweek/comcast-one/server/search"
)

func init() {
	search.OpenDB()
}

func main() {
	// search.SearchForQuery("conference")
	// search.ReadDBStream()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handler.SearchHandler).Methods("GET")
	muxRouter.HandleFunc("/bookmarks", handler.SearchBookmarksHandler).Methods("GET")
	muxRouter.HandleFunc("/create", handler.CreateBookmark).Methods("POST")
	muxRouter.HandleFunc("/update", handler.UpdateHandler).Methods("POST")
	server := &http.Server{
		Handler:      handlers.CORS(headersOk, originsOk, methodsOk)(muxRouter),
		Addr:         "0.0.0.0:9000",
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
	}
	log.Fatal(server.ListenAndServe())

}
