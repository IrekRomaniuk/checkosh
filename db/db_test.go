package db

import (
	"testing"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func TestUpdate(t *testing.T) {	
	robo := map[string]string{"Name": "Irek-11", "Ext": "1.1.1.1", "Policy": "TestUpdate", "Int":"10.10.10.10"}
	session, _ := Connect("10.254.253.100:27017")
	err := Update(robo, session, "meteor", "robo", "Name")	
	if err != nil { fmt.Println(robo, err)}
	robo = map[string]string{"Name": "Irek-11", "Ext": "1.1.1.1", "Policy": "TestUpdate-see Int present"}
	if err != nil { fmt.Println(robo, err)}
}

func TestUpsert(t *testing.T) {
	robo := map[string]string{"Name": "Irek-11", "Ext": "2.2.2.2", "Policy": "TestUpsert-see Int removed"}
	session, _ := Connect("10.254.253.100:27017")
	_, err := Upsert(robo, session, "meteor", "robo", "Name")
	if err != nil { fmt.Println(robo, err)}
}

/*func TestStats(t *testing.T) {
	var results []string // map[string]string
	session, _ := Connect("10.254.253.100:27017")
	c := session.DB("meteor").C("robo")
	_ = c.Find(bson.M{"Name": "Irek-11"}).Sort("-timestamp").All(&results)
	fmt.Println("All Irek-11: ", results)
}*/