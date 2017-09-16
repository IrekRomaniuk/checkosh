package main

import (
	"github.com/IrekRomaniuk/checkosh/db"
	"log"
	"flag"
	"fmt"
	"os"
)

var (
	//PATH to results
	PATH  = flag.String("p", "./", "path to read files from")
	//ADDRESS of db
	ADDRESS  = flag.String("a", "10.254.253.100:27017", "db addr")
	version   = flag.Bool("v", false, "Prints current version")
	// Version : Program version
	Version   = "No Version Provided" 
	// BuildTime : Program build time
	BuildTime = ""
)

const (
	database = "meteor"
	collection = "robo"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2017 @IrekRomaniuk. All rights reversed.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *version {
		fmt.Printf("App Version: %s\nBuild Time : %s\n", Version, BuildTime)
		os.Exit(0)
	}	
}

func main() {
	log.Println("checkosh - starting Read ")
	// Map first column of f file to db Name and 3rd to db Ext, Policy always mapped to last column
	f := db.File{*PATH + "checkosh-lsm", map[string]int{"Name":0, "Ext":2, "Policy":0}}
	err := db.ReadFile(*ADDRESS, database, collection, f, true) //true for upsert
	if err != nil { 
		log.Println(err) 
	} else {
		log.Println("checkosh - 1st Read completed ")
	}

	f = db.File{*PATH + "./checkosh-int", map[string]int{"Name":0, "Int":1, "Comment":2}}
	err = db.ReadFile("10.254.253.100:27017", database, collection, f, false) //false if update
	if err != nil { 
		log.Println(err) 
	} else {
		log.Println("checkosh - 2nd Read completed ")
	}
}
