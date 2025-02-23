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
		panic(err)
	}

	filter := bson.D{{Key: "_id", Value: idParam}}
	log.Println(filter)
	var result Movie
	err = coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Movie{}, errors.Wrap(err, "no movie found")
		}
		panic(err)
	}

	return result, nil
}

func (r MovieRepository) Create(movie Movie) (string, error) {
	coll := r.dbclient.DB().Database("sample_mflix").Collection("movies")

	result, err := coll.InsertOne(context.Background(), movie)

	if err != nil {
		return "", errors.Wrap(err, "failed to insert movie")
	}
	log.Println(result)
	return result.InsertedID.(primitive.ObjectID).Hex(), nil

	// fmt.Println("Inserted document ID:", result.InsertedID.(primitive.ObjectID).Hex())
}
