package repository

import (
	"context"
	"log"

	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Movie struct {
	ID    primitive.ObjectID `bson:"_id"`
	Title string
}

type IMovieRepository interface {
	FindById(id string) (Movie, error)
	Create(movie Movie) (string, error)
	Update(id string, movie Movie) error
	Delete(id string) error
}

type MovieRepository struct {
	dbclient database.IdbService
}

func NewMovieRepository(db database.IdbService) *MovieRepository {
	return &MovieRepository{
		dbclient: db,
	}
}

func (r MovieRepository) FindById(id string) (Movie, error) {
	coll := r.dbclient.DB().Database("sample_mflix").Collection("movies")

	idParam, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return Movie{}, errors.Wrap(err, "not valid id")
	}

	filter := bson.D{{Key: "_id", Value: idParam}}
	log.Println(filter)
	var result Movie
	err = coll.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Movie{}, errors.Wrap(err, "not movie found")
		}
		return Movie{}, errors.Wrap(err, "error finding movie")
	}

	return result, nil
}

func (r MovieRepository) Create(movie Movie) (string, error) {
	coll := r.dbclient.DB().Database("sample_mflix").Collection("movies")

	result, err := coll.InsertOne(context.Background(), movie)

	if err != nil {
		return "", errors.Wrap(err, "failed to insert movie")
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r MovieRepository) Update(id string, movie Movie) error {
	coll := r.dbclient.DB().Database("sample_mflix").Collection("movies")

	idParam, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errors.Wrap(err, "not valid id")
	}

	filter := bson.D{{Key: "_id", Value: idParam}}

	_, err = coll.ReplaceOne(context.Background(), filter, movie)

	if err != nil {
		return errors.Wrap(err, "failed to update movie")
	}

	return nil
}

func (r MovieRepository) Delete(id string) error {
	coll := r.dbclient.DB().Database("sample_mflix").Collection("movies")

	idParam, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errors.Wrap(err, "not valid id")
	}

	filter := bson.D{{Key: "_id", Value: idParam}}

	_, err = coll.DeleteOne(context.Background(), filter)

	if err != nil {
		return errors.Wrap(err, "failed to delete movie")
	}

	return nil
}
