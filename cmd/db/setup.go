package main

import (
	"github.com/globalsign/mgo"
	"log"
	"os"
	"time"
)

func main() {
	var err error
	mongoHost := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB")
	userName := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASS")
	if mongoHost == "" || userName == "" || password == "" {
		log.Fatal("Mongo info not provided")
	}

	mongoDialInfo := &mgo.DialInfo{
		Addrs: []string{mongoHost},
		Database: dbName,
		Username: userName,
		Password: password,
		Timeout: 10 * time.Second,
	}
	session, err := mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		log.Fatalf("Unable to create Mongo session: %s", err.Error())
	}
	defer session.Close()

	err = session.DB("").AddUser("mongo_user", "mongo_secret", false)

	info := &mgo.CollectionInfo{}
	err = session.DB("").C("kudos").Create(info)

	if err != nil {
		log.Fatal(err)
	}
}
