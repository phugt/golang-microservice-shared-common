package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email      string             `bson:"email,omitempty" json:"email"`
	Password   string             `bson:"password,omitempty" json:"-"`
	FullName   string             `bson:"fullName,omitempty" json:"fullName"`
	Address    string             `bson:"address,omitempty" json:"address"`
	Desc       string             `bson:"desc,omitempty" json:"desc"`
	JoinTime   int64              `bson:"joinTime,omitempty"  json:"joinTime"`
	DeleteTime int64              `bson:"deleteTime,omitempty"  json:"deleteTime"`
}
