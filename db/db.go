package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"bufio"
	"os"
	"strings"
	//"log"
)
// Session to db
type Session *mgo.Session
// Robo struct
type Robo struct {
	Name, Ext, Int, Policy, DC1, DC2 string
}
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

// Upsert inserts document into collection
func Upsert(robo Robo, session *mgo.Session, db, collection string, Name string) (*mgo.ChangeInfo, error) {
	c := session.DB(db).C(collection)
	info, err := c.Upsert(bson.M{"name": Name}, robo)
	return info, err		
  }
// ReadFile reads file into db
func ReadFile(address, database, collection string, f File) error {
	file, err := os.Open(f.Name)	
	if err != nil { return err }	
	defer file.Close()
	session, err := Connect(address)
	if err != nil { return err }
	
	// create a new scanner and read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		robo := strings.Fields(scanner.Text())
		//log.Println(robo[f.Mapping["Name"]], robo[f.Mapping["Ext"]], robo[len(robo)-1])	
		r := Robo{Name:robo[f.Mapping["Name"]], Ext:robo[f.Mapping["Ext"]], 
			Int:robo[f.Mapping["Int"]], Policy:robo[len(robo)-1]}	
		_, err = Upsert(r, session, database, collection, robo[0])
		if err != nil { return err }
	}
	err = scanner.Err() 
	return err	  
}

// Insert inserts document into collection
 func Insert(robo Robo, session *mgo.Session, db, collection string) error {
	c := session.DB(db).C(collection)
	err := c.Insert(robo)
	return err		
  }

// Close disconnects from db
func Close(session *mgo.Session) {
	session.Close()
  }