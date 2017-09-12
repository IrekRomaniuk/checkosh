package main

import (
	"os"
	"bufio"
	"github.com/IrekRomaniuk/checkosh/db"
	"strings"
)

func main() {
	// Record in db
	session, err := db.Connect("10.254.253.100:27017")
	if err != nil { os.Exit(2) }
	err = readFile(session, "checkosh1")
	if err != nil { os.Exit(1) }
	db.Close(session)
}

func readFile(session db.Session, f string) error {
	file, err := os.Open(f)	
	if err != nil { return err }	
	defer file.Close()
	// create a new scanner and read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		robo := strings.Fields(scanner.Text())
		//fmt.Println(robo[0], robo[2], robo[len(robo)-1])
		r := db.Robo{Name:robo[0], Ext:robo[2], Policy:robo[len(robo)-1] }		
		_, err = db.Upsert(r, session, "meteor", "robo", robo[0])
		if err != nil { return err }
	}
	err = scanner.Err() 
	return err	  
}
