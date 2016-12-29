package logoperation

import "gopkg.in/mgo.v2-unstable"

type logserver struct {
	session *mgo.Session
}

type M map[string]interface{}

type Log struct {
	Title      string
	Content    M
	Level      string
	Createdate string
}

//NewLogServer ...
func NewLogServer(host string) (*logserver, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	return &logserver{session}, nil
}
