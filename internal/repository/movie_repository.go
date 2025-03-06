package repository

import (
	"context"
	"time"

	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/melnikdev/go-grafana/internal/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMovieRepository interface {
	GetTopMovies(limit int64) ([]model.Movie, error)
	FindById(id string) (model.Movie, error)
	Create(movie model.Movie) (string, error)
	Update(id string, movie model.Movie) error
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

func (r MovieRepository) GetTopMovies(limit int64) ([]model.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := r.dbclient.DB().Database("sample_mflix").Collection("movies")

	filter := bson.D{{Key: "imdb.rating", Value: bson.D{{Key: "$ne", Value: nil}}}, {Key: "poster", Value: bson.D{{Key: "$ne", Value: nil}}}}
	options := options.Find().SetSort(bson.D{{Key: "imdb.rating", Value: -1}}).SetLimit(limit)

	cursor, err := coll.Find(ctx, filter, options)

	if err != nil {
		return nil, errors.Wrap(err, "error getting top 5 movies")
	}

	var movies []model.Movie
	err = cursor.All(context.Background(), &movies)

	if err != nil {
		return nil, errors.Wrap(err, "error getting movies")
	}

	return movies, nil

}

func (r MovieRepository) FindById(id string) (model.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := r.dbclient.DB().Database("sample_mflix").Collection("movies")

	idParam, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return model.Movie{}, errors.Wrap(err, "not valid id")
	}

	filter := bson.D{{Key: "_id", Value: idParam}}

	var result model.Movie
	err = coll.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Movie{}, errors.Wrap(err, "not movie found")
		}
		return model.Movie{}, errors.Wrap(err, "error finding movie")
	}

	return result, nil
}

func (r MovieRepository) Create(movie model.Movie) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := r.dbclient.DB().Database("sample_mflix").Collection("movies")

	result, err := coll.InsertOne(ctx, movie)

	if err != nil {
		return "", errors.Wrap(err, "failed to insert movie")
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r MovieRepository) Update(id string, movie model.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := r.dbclient.DB().Database("sample_mflix").Collection("movies")

	idParam, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errors.Wrap(err, "not valid id")
	}

	filter := bson.D{{Key: "_id", Value: idParam}}

	_, err = coll.ReplaceOne(ctx, filter, movie)

	if err != nil {
		return errors.Wrap(err, "failed to update movie")
	}

	return nil
}

func (r MovieRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := r.dbclient.DB().Database("sample_mflix").Collection("movies")

	idParam, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errors.Wrap(err, "not valid id")
	}

	filter := bson.D{{Key: "_id", Value: idParam}}

	_, err = coll.DeleteOne(ctx, filter)

	if err != nil {
		return errors.Wrap(err, "failed to delete movie")
	}

	return nil
}
