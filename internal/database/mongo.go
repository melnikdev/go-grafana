package database

import (
	"context"
	"log"
	"time"

	"github.com/melnikdev/go-grafana/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IdbService interface {
	Health() map[string]string
	Disconnect(ctx context.Context) error
	DB() *mongo.Client
}

type dbService struct {
	db *mongo.Client
}

func New(config *config.MongoDB) IdbService {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(config.Uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return &dbService{
		db: client,
	}
}

func (s *dbService) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("db down: %v", err)
	}

	log.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *dbService) Disconnect(ctx context.Context) error {
	return s.db.Disconnect(ctx)
}

func (s *dbService) DB() *mongo.Client {
	return s.db
}
