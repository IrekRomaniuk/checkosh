package main

import (
	"github.com/IrekRomaniuk/checkosh/db"
	"log"
	"time"
)

const (
	database = "meteor"
	collection = "robo"
)

func main() {
	// Map first column of f file to db Name and 3rd to db Ext, Policy always mapped to last column
	f := db.File{"./data/checkosh-lsm", map[string]int{"Name":0, "Ext":2, "Policy":0}}
	err := db.ReadFile("10.254.253.100:27017", database, collection, f, true) //true for upsert
	if err != nil { 
		log.Println(err) 
	} else {
		log.Printf("1st Read completed %t", time.Now())
	}

	f = db.File{"./data/checkosh-int", map[string]int{"Name":0, "Int":1, "Comment":2}}
	err = db.ReadFile("10.254.253.100:27017", database, collection, f, false) //false if update
	if err != nil { 
		log.Println(err) 
	} else {
		log.Printf("2nd Read completed %t", time.Now())
	}
}
