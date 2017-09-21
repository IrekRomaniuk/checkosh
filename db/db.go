package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"bufio"
	"os"
	"strings"
	"fmt"
	"path/filepath"
)
// Session to db
type Session *mgo.Session
// File struct
type File struct {
	Name string
	Mapping map[string]int
}
// Connect to db
func Connect(server  string) (*mgo.Session, error) {
	//fmt.Println("Connecting to db")
	  session, err := mgo.Dial(server)
	if err != nil {
		  return nil, err
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	return session, nil  
  }
// Upsert inserts document into collection - robo Robo
func Upsert(robo map[string]string, session *mgo.Session, db, collection, Name string) (*mgo.ChangeInfo, error) {
	c := session.DB(db).C(collection)
	info, err := c.Upsert(bson.M{"Name": robo["Name"]}, robo)
	return info, err		
	}
// Update updates document 
func Update(robo map[string]string, session *mgo.Session, db, collection, Name string) error {
	c := session.DB(db).C(collection)
	err := c.Update(bson.M{"Name": robo["Name"]}, bson.M{"$set": robo})
	return err		
	}
// var r Robo
var robo []string
// ReadFile reads file into db
func ReadFile(address, database, collection string, f File) error { //, upsert bool
	file, err := os.Open(f.Name)	
	if err != nil { return err }	
	defer file.Close()
	//fmt.Println(filepath.Dir(f.Name) + "/bin/IP2LOCATION-LITE-DB11.BIN")
	session, err := Connect(address)
	if err != nil { return err }
	// create a new scanner and read the file line by line
	scanner := bufio.NewScanner(file)
	r := make(map[string]string)
	for scanner.Scan() {
		robo = strings.Fields(scanner.Text())
		val, ok := f.Mapping["Name"]
		if ok { r["Name"] = robo[val] }
		val, ok = f.Mapping["Ext"]
		if ok { 
			r["Ext"] = robo[val]
			record, _ := Location(filepath.Dir(f.Name) + "/bin/IP2LOCATION-LITE-DB11.BIN", r["Ext"]) 
			r["City"] = record.City
			r["Latitude"] = fmt.Sprintf("%.6f", record.Latitude)
			r["Longitude"] = fmt.Sprintf("%.6f", record.Longitude)
			r["Region"] = record.Region
			r["Zipcode"] = record.Zipcode
		}
		val, ok = f.Mapping["Int"]
		if ok { r["Int"] = robo[val] }
		val, ok = f.Mapping["Comment"]
		if ok { r["Comment"] = strings.Join(robo[val:],"") }
		_, ok = f.Mapping["Policy"]
		if ok { r["Policy"] = robo[len(robo)-1] } 
		if err = Update(r, session, database, collection, r["Name"]); err != nil {
			_, err = Upsert(r, session, database, collection, r["Name"])
		}
		if err != nil { 
			fmt.Println(r["Name"], err)
			return err 
		}
	}
	err = scanner.Err() 
	return err	  
}
// Insert inserts document into collection
 func Insert(robo map[string]string, session *mgo.Session, db, collection string) error {
	c := session.DB(db).C(collection)
	err := c.Insert(robo)
	return err		
  }
// Close disconnects from db
func Close(session *mgo.Session) {
	session.Close()
	}
// Intitialize initis db, see also 'https://gist.github.com/border/3489566'
func Intitialize(address, database, collection string) error {
	session, err := Connect(address)
	if err != nil { return err }
	err = session.DB(database).DropDatabase()
	if err != nil { return err }
	c := session.DB(database).C(collection)
	// Index
	index := mgo.Index{
		Key:        []string{"Int", "Name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		return(err)
	}
	// Fill in Int with 10.192.0.0/12
  var ip string
	for x := 192; x <= 207; x++ {
		for y := 0; y <= 255; y++ {
			ip = fmt.Sprintf("10.%d.%d.1", x, y)
			_, err = c.Upsert(bson.M{"_id": ip}, map[string]string{})
			if err != nil { fmt.Println(ip, err)}	else {
				fmt.Println(ip)	
			}	
		}
	}
	return err
}