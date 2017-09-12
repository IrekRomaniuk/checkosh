package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
// Session to db
type Session *mgo.Session
// Robo struct
type Robo struct {
	Name, Ext, Int, Policy, DC1, DC2 string
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