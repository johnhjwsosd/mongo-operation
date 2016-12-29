package logoperation

import (
	mgo "gopkg.in/mgo.v2-unstable"
	"gopkg.in/mgo.v2-unstable/bson"
)

func (logserver *logserver) Writelog(db, collections string, log *Log) (err error) {
	logserver.session.SetMode(mgo.Monotonic, true)
	c := logserver.session.DB(db).C(collections)
	err = c.Insert(log)
	return
}

func (logserver *logserver) ReadlogALL(db, collections string, where ...bson.M) (*[]Log, error) {
	logserver.session.SetMode(mgo.Monotonic, true)
	c := logserver.session.DB(db).C(collections)
	result := []Log{}
	m := bson.M{}
	if len(where) > 0 {
		m = where[0]
	}
	err := c.Find(m).All(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
