package models

import "time"

type Todo struct {
	ID        int       `json:"id" bson:"_id"`
	Data      string    `json:"data"  bson:"data"`
	IsDone    bool      `json:"isDone" bson:"isDone"`
	CreatedAt time.Time `json:"create_at" bson:"create_at"`
	IsStared  bool      `json:"isStared"  bson:"isStared"`
}
