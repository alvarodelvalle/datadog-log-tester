package main

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type pii struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	SecretKey      string `json:"secret_key"`
	SecretPassword string `json:"secret_password"`
	IpAddress      string `json:"ip_address"`
	CcNumber       string `json:"cc_number"`
	PostalCode     int    `json:"postal_code"`
	Ssn            string `json:"ssn"`
}

func main() {
	p1 := &pii{
		Name:           "Vars Gardian",
		Email:          "vars@geemail.com",
		SecretKey:      "3352dgad#KDJSIEQ",
		SecretPassword: "dontsharethispassword",
		IpAddress:      "10.10.12.1",
		CcNumber:       "4111-1111-1111-1111",
		PostalCode:     33545,
		Ssn:            "111-11-1111",
	}

	// use JSONFormatter
	log.SetFormatter(&log.JSONFormatter{})
	// log an event as usual with logrus
	for i := 0; i < 20; i++ {
		log.WithFields(log.Fields{"user":p1}).Info("User Data")
		i++
		time.Sleep(5 * time.Second)
	}
}