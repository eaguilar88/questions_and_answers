package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

//Question _
type Question struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Answers     []*Answer          `bson:"answers" json:"answers"`
}

//Questions _
type Questions []Question

//Answer _
type Answer struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Answer   string             `bson:"answer" json:"answer"`
}
