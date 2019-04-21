package mongo

import (
	"crypto/tls"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/globalsign/mgo"
)

const (
	KeyMongodbServer         = "MONGODB_SERVER"
	KeyMongodbDatabase       = "MONGODB_DATABASE"
	KeyMongodbUsername       = "MONGODB_USERNAME"
	KeyMongodbPassword       = "MONGODB_PASSWORD"
	KeyMongodbReplicaSetName = "MONGODB_REPLICASET_NAME"
	KeyMongodbSSL            = "MONGODB_SSL"
)

func getDialInfo() *mgo.DialInfo {
	dbHost := strings.Split(os.Getenv(KeyMongodbServer), ",")
	dbName := os.Getenv(KeyMongodbDatabase)
	dbUser := os.Getenv(KeyMongodbUsername)
	dbPassword := os.Getenv(KeyMongodbPassword)
	replicaSetName := os.Getenv(KeyMongodbReplicaSetName)

	log.Println("MongodbServer: " + os.Getenv(KeyMongodbServer))
	log.Println("DBName: " + dbName)
	log.Println("SSL: " + os.Getenv(KeyMongodbSSL))

	dialInfo := &mgo.DialInfo{
		Addrs:          dbHost,
		Database:       dbName,
		Username:       dbUser,
		Password:       dbPassword,
		ReplicaSetName: replicaSetName,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{InsecureSkipVerify: true})
		},
		Timeout: time.Second * 10,
	}

	if os.Getenv(KeyMongodbSSL) != "true" {
		dialInfo.DialServer = nil
	}

	return dialInfo
}
