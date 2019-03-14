package entity

import "time"

type History struct {
	Id       string    `bson:"_id"`
	Name     string    `bson:"name"`
	Filename string    `bson:"filename"`
	Info     string    `bson:"info"`
	Time     time.Time `bson:"time"`
}
