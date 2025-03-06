package repository

import (
	"context"
	"time"

	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/melnikdev/go-grafana/internal/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserRepository interface {
	Create(user model.User) (string, error)
}

type UserRepository struct {
	dbclient database.IdbService
}

func NewUserRepository(db database.IdbService) *UserRepository {
	return &UserRepository{
		dbclient: db,
	}
}

func (r UserRepository) Create(user model.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := r.dbclient.DB().Database("movies_test_db").Collection("users")

	result, err := coll.InsertOne(ctx, user)

	if err != nil {
		return "", errors.Wrap(err, "failed to insert user")
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
