//go:build mongodb

package geo

import (
	log "github.com/sirupsen/logrus"

	"github.com/ucy-coast/hotel-app/pkg/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DatabaseSession struct {
	MongoSession *mgo.Session
}

func NewDatabaseSession(db_addr string) *DatabaseSession {
	// TODO: Implement me
	session, err := mgo.Dial(db_addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("New session successfull...")

	return &DatabaseSession{
		MongoSession: session,
	}
}

func (db *DatabaseSession) LoadDataFromJsonFile(rateJsonPath string) {
	util.LoadDataFromJsonFile(db.MongoSession, "geo-db", "geo", rateJsonPath)
}

// newGeoIndex returns a geo index with points loaded
func (db *DatabaseSession) newGeoIndex() *geoindex.ClusteringIndex {
	// add points to index
	index := geoindex.NewClusteringIndex()


	return index
}
