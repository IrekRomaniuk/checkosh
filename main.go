package main

import (
	"os"
	"github.com/IrekRomaniuk/checkosh/db"
)

const (
	database = "meteor"
	collection = "robo"
)

func main() {
	f := db.File{"./data/checkosh-lsm", map[string]int{"Name":0, "Ext":2}}
	err := db.ReadFile("10.254.253.100:27017", database, collection, f)
	if err != nil { os.Exit(1) }
}


