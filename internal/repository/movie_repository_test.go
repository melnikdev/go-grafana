package repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/melnikdev/go-grafana/internal/config"
	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/melnikdev/go-grafana/internal/model"
	"github.com/melnikdev/go-grafana/internal/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var testDB database.IdbService
var testRepo *repository.MovieRepository

func setupTestDB(t *testing.T) {
	if testDB == nil {
		mongoURI := os.Getenv("MONGO_URI")
		if mongoURI == "" {
			mongoURI = "mongodb://localhost:27017"
		}
		testDB = database.New(&config.MongoDB{Uri: mongoURI})
	}

	clearCollection(t)
	testRepo = repository.NewMovieRepository(testDB)
}

func clearCollection(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := testDB.DB().Database("sample_mflix").Collection("movies")
	_, err := collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		t.Fatalf("Failed to clear test database: %v", err)
	}
}

func TestCreateMovie(t *testing.T) {
	setupTestDB(t)

	movie := model.Movie{
		ID:     primitive.NewObjectID(),
		Title:  "Test Movie",
		Plot:   "This is a test movie",
		Poster: "test_poster.jpg",
		Imdb:   model.Imdb{Rating: "8.5", Votes: "1000", Id: 123456},
	}

	id, err := testRepo.Create(movie)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestFindByIdMovie(t *testing.T) {
	setupTestDB(t)

	movie := model.Movie{
		ID:     primitive.NewObjectID(),
		Title:  "Find Me",
		Plot:   "This is a test movie",
		Poster: "test_poster.jpg",
		Imdb:   model.Imdb{Rating: "7.0", Votes: "500", Id: 654321},
	}

	id, _ := testRepo.Create(movie)

	fetchedMovie, err := testRepo.FindById(id)
	assert.NoError(t, err)
	assert.Equal(t, movie.Title, fetchedMovie.Title)
}

func TestGetTopMovies(t *testing.T) {
	setupTestDB(t)

	movie1 := model.Movie{
		ID:     primitive.NewObjectID(),
		Title:  "Top Movie 1",
		Plot:   "Movie 1",
		Poster: "poster1.jpg",
		Imdb:   model.Imdb{Rating: "9.0", Votes: "10000", Id: 1},
	}
	movie2 := model.Movie{
		ID:     primitive.NewObjectID(),
		Title:  "Top Movie 2",
		Plot:   "Movie 2",
		Poster: "poster2.jpg",
		Imdb:   model.Imdb{Rating: "8.5", Votes: "8000", Id: 2},
	}

	testRepo.Create(movie1)
	testRepo.Create(movie2)

	topMovies, err := testRepo.GetTopMovies(2)
	assert.NoError(t, err)
	assert.Len(t, topMovies, 2)
	assert.Equal(t, "Top Movie 1", topMovies[0].Title)
}

func TestUpdateMovie(t *testing.T) {
	setupTestDB(t)

	movie := model.Movie{
		ID:     primitive.NewObjectID(),
		Title:  "Original Title",
		Plot:   "This is a test movie",
		Poster: "test_poster.jpg",
		Imdb:   model.Imdb{Rating: "8.0", Votes: "2000", Id: 789123},
	}

	id, _ := testRepo.Create(movie)

	updatedMovie := movie
	updatedMovie.Title = "Updated Title"

	err := testRepo.Update(id, &updatedMovie)
	assert.NoError(t, err)

	fetchedMovie, _ := testRepo.FindById(id)
	assert.Equal(t, "Updated Title", fetchedMovie.Title)
}

func TestDeleteMovie(t *testing.T) {
	setupTestDB(t)

	movie := model.Movie{
		ID:     primitive.NewObjectID(),
		Title:  "Delete Me",
		Plot:   "This movie will be deleted",
		Poster: "delete_poster.jpg",
		Imdb:   model.Imdb{Rating: "5.0", Votes: "500", Id: 999999},
	}

	id, _ := testRepo.Create(movie)

	err := testRepo.Delete(id)
	assert.NoError(t, err)

	_, err = testRepo.FindById(id)
	assert.Error(t, err)
}
