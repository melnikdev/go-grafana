package repository

import (
	"context"

	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Restaurant struct {
	ID    primitive.ObjectID `bson:"_id"`
	Title string
}

type IMovieRepository interface {
	FindById() (Restaurant, error)
}

type MovieRepository struct {
	dbclient database.IdbService
}

func NewMovieRepository(db database.IdbService) *MovieRepository {
	return &MovieRepository{
		dbclient: db,
	}
}

func (r MovieRepository) FindById() (Restaurant, error) {
	coll := r.dbclient.DB().Database("sample_mflix").Collection("movies")

	filter := bson.D{{Key: "title", Value: "The Great Train Robbery"}}

	var result Restaurant
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Restaurant{}, errors.New("no documents found")
		}
		panic(err)
	}

	return result, nil
}
