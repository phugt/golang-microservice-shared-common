package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserToken struct {
	ID         string             `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"userId,omitempty" json:"userId"`
	UserAgent  string             `bson:"userAgent,omitempty" json:"userAgent"`
	CreateTime int64              `bson:"joinTime,omitempty" json:"createTime"`
}
