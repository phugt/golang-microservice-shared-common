package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLog struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID primitive.ObjectID `bson:"userId,omitempty" json:"userId"`
	Action string             `bson:"action,omitempty" json:"action"`
	Method string             `bson:"method,omitempty" json:"method"`
	Data   string             `bson:"data,omitempty" json:"data"`
	Time   int64              `bson:"time,omitempty" json:"time"`
}
