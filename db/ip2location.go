package db

import (	
	"os"
	"github.com/ip2location/ip2location-go"
)
// Location return lat and long
func Location(file, ip string) (ip2location.IP2Locationrecord, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return ip2location.IP2Locationrecord{}, err
	  }

	ip2location.Open(file)

	record := ip2location.Get_all(ip)
		
	ip2location.Close()

	return record, nil
}