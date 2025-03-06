package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Plot   string             `bson:"plot,omitempty"`
	Poster string             `bson:"poster,omitempty"`
	Imdb   Imdb               `bson:"imdb,omitempty"`
}

type Imdb struct {
	Rating string `bson:"rating,omitempty"`
	Votes  string `bson:"votes,omitempty"`
	Id     int32  `bson:"id,omitempty"`
}
