package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"userId,omitempty" json:"userId"`
	User       User               `json:"user"`
	Roles      []string           `bson:"roles" json:"roles"`
	JoinTime   int64              `bson:"joinTime,omitempty" json:"joinTime"`
	DeleteTime int64              `bson:"deleteTime,omitempty" json:"deleteTime"`
}
