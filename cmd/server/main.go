package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/joho/godotenv"
	"github.com/teste-transfeera/internal/graph"
	"github.com/teste-transfeera/internal/repository"
	"github.com/teste-transfeera/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctx := context.Background() //context.WithTimeout(context.Background(), 5*time.Second)

	collection := initDB(ctx)
	receiverRepository := repository.NewReceiverRepository(collection, ctx)

	receiverUsecases := usecase.NewReceiverUseCases(receiverRepository)

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ReceiverUseCases: receiverUsecases}}))
	http.Handle("/api/v1", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initDB(ctx context.Context) *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URL"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("transfeera").Collection("receiver")
}
