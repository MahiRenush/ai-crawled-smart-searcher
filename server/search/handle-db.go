package search

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imageurl"`
	Website     string    `json:"website"`
	Upvotes     int       `json:"upvotes"`
	Comments    int       `json:"comments"`
	Updated     time.Time `json:"updated"`
}

type Bookmark struct {
	Title string
	Url   string
}

var err, anerr error
var bucketName = []byte("evtBucket")
var db *bolt.DB
var bucketName1 = []byte("bookmarks")

// Field name is not needed until read of one Event i.e., ReadDB()
var fieldName = []byte("eventsKey3")

func OpenDB() {
	db, err = bolt.Open("event.db", 0600, nil)
	if err != nil {
		fmt.Println("Open ERROR:", err)
	}
	fmt.Println(db.GoString(), db.Info())

}

func createEvent(events *Event) {
	anerr = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			fmt.Println("ERROR on Bucket creation if not exists: ", err)
		}

		// Generate ID for the events.
		id, _ := b.NextSequence()
		events.ID = int(id)
		fmt.Println(id, events.ID)

		// Marshal events data into bytes.
		buf, err := json.Marshal(events)
		if err != nil {
			fmt.Println("Marshall ERROR: ", err)
		}

		// Persist bytes to users bucket.
		err = b.Put(itob(events.ID), buf)

		// Additional Get for checking the field added
		/*
			v := b.Get(itob(events.ID))
			fmt.Printf("The answer is: %s\n", v)
		*/
		return err
	})
	if anerr != nil {
		fmt.Println("Update Error:", anerr)
	}
}

// add bookmarks

func AddBookmarks(Title string, Url string) {

	if db.Path() == "" {
		OpenDB()
	}
	var bookmark Bookmark
	bookmark.Title = Title
	bookmark.Url = Url
	anerr = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucketName1)
		if err != nil {
			fmt.Println("ERROR on Bucket creation if not exists: ", err)
		}

		// Generate ID for the events.
		id, _ := b.NextSequence()
		key := int(id)
		fmt.Println(id, key)

		buf, err := json.Marshal(bookmark)
		if err != nil {
			fmt.Println("Marshall ERROR: ", err)
		}
		err = b.Put(itob(key), buf)

		return err
	})
	if anerr != nil {
		fmt.Println("Update Error:", anerr)
	}

}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi converts from byte to int
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func ReadDB() {
	anerr = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		v := b.Get(fieldName)
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
	if anerr != nil {
		fmt.Println("View DB ERROR:", anerr)
	}
}

func ReadBookmarks() []Bookmark {

	if db.Path() == "" {
		OpenDB()
	}
	var m []Bookmark
	var s *Bookmark
	var b *bolt.Bucket
	var c *bolt.Cursor
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b = tx.Bucket(bucketName1)
		if b != nil {
			c = b.Cursor()
		}

		if c != nil {
			fmt.Println("c: ", c)
			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("key=%d, value=%s\n", btoi(k), v)
				json.Unmarshal(v, &s)
				m = append(m, *s)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return m
}

func ReadDBStream() []Event {

	if db.Path() == "" {
		OpenDB()
	}
	var m []Event
	var s *Event
	var b *bolt.Bucket
	var c *bolt.Cursor
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b = tx.Bucket(bucketName)
		if b != nil {
			c = b.Cursor()
		}

		if c != nil {
			fmt.Println("c: ", c)
			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("key=%d, value=%s\n", btoi(k), v)
				json.Unmarshal(v, &s)
				m = append(m, *s)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return m
}

func WriteDB() {
	events := []Event{
		{1, "Product Experience", "A website to learn more about our products and services", "https://www.comcasttechnologysolutions.com/sites/default/files/styles/resource_teaser/public/2022-08/iStock-1129543888.jpg?itok=nvkylwvP", "https://px.comcast.com/", 15, 12, time.Date(2016, 7, 13, 23, 59, 0, 0, time.UTC)},
		{2, "Comcast Site", "A website for comcast", "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/Comcast_Logo.svg/2560px-Comcast_Logo.svg.png", "https://www.comcasttechnologysolutions.com/", 13, 6, time.Date(2016, 7, 13, 23, 59, 0, 0, time.UTC)},
		{3, "Workday", "Compensation, Benefits, Performance, etc", "https://www.workday.com/content/dam/web/en-us/images/social/workday-og-theme.png", "https://wd5.myworkday.com/comcast/d/home.htmld", 2, 5, time.Date(2016, 2, 20, 23, 59, 0, 0, time.UTC)},
		{4, "CVP Console", "Cloud Video Platform is a place where all Video related products like media works. Also this CVP loom application is the modern one that succeds deprecated flash console (adobe flash plugin) and is made of HTML5.", "https://i.ibb.co/PjrPrd8/cvp.png", "http://console.theplatform.com/", 3, 6, time.Date(2016, 7, 13, 23, 59, 0, 0, time.UTC)},
		{5, "Udemy - Comcast", "The udemy learning website by Comcast", "https://findlogovector.com/wp-content/uploads/2018/11/udemy-logo-vector.png", "https://comcast.udemy.com/", 123, 34, time.Date(2016, 7, 13, 23, 59, 0, 0, time.UTC)},
		{6, "Comcast Now", "The Goto website for Comcast named comcast now", "https://www.comcasttechnologysolutions.com//themes/custom/themekit/images/logo-inverse.png", "http://www.comcastnow.com/", 2, 5, time.Date(2016, 2, 20, 23, 59, 0, 0, time.UTC)},
		{7, "LRM - Prod", "LRM - Linear Right Metadata Management Prod URL", "https://lrmui.aort.theplatform.com/assets/img/lrm_logov2.svg", "https://lrmui.theplatform.com/", 5, 3, time.Date(2016, 7, 11, 0, 0, 0, 0, time.UTC)},
		{8, "ADP Portal", "ADP offers industry-leading online payroll software & HR services.", "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSFyxzH12YgcozYBzT3WdztPdoyoMBG-BHuHdZ3iJ8EgjBDvisdBm_19Q1l8XdraXQEjBc&usqp=CAU", "https://www.vista.adp.com/ESS4/ESSV5/login", 10, 15, time.Date(2016, 7, 13, 23, 59, 0, 0, time.UTC)},
		{9, "Docs - The Platform", "The documents of the products and services provided by the platform is found here", "https://www.comcasttechnologysolutions.com//themes/custom/themekit/images/logo-inverse.png", "https://docs.theplatform.com/", 4, 5, time.Date(2016, 7, 13, 23, 59, 0, 0, time.UTC)},
		{10, "CIEC ESS Portal", "Employee service portal that helps us in all ways and have collection of website links needed for comcast employee.", "https://www.comcasttechnologysolutions.com//themes/custom/themekit/images/logo-inverse.png", "https://ciecessportal.comcast.net/", 2, 3, time.Date(2016, 7, 13, 23, 59, 0, 0, time.UTC)},
		{11, "Business - Comcast", "", "https://cdn.pdc.business.comcast.com/~/media/business_comcast_com/images/Re-Arch/Internet/WifiPro%20SpotlightTopImage_01.png?rev=8a55455c-22fc-4099-af67-a6eb7f8907d9&h=245&w=433&la=en&hash=A4586F88E12A84C504205F8225619B6568AD3BE5", "https://www.comcasttechnologysolutions.com/", 1, 0, time.Date(2016, 7, 13, 23, 59, 0, 0, time.UTC)},
		{12, "Company - Comcast", "", "https://www.cmcsa.com/sites/g/files/knoqqb64786/themes/site/nir_pid1513/dist/images/hero/ir.jpg", "https://www.cmcsa.com/", 28, 3, time.Date(2016, 7, 13, 23, 59, 0, 0, time.UTC)},
	}
	for i := 0; i < len(events); i++ {
		createEvent(&events[i])
	}
}
func CloseDB() {
	db.Close()
}
