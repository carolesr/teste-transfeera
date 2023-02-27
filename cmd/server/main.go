package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
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
		port = ":8080"
	}

	ctx := context.Background()

	collection := initDB(ctx)
	receiverRepository := repository.NewReceiverRepository(collection, ctx)

	receiverUsecases := usecase.NewReceiverUseCases(receiverRepository)

	initServer(port, receiverUsecases)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil && err != http.ErrServerClosed {
			log.Fatal("error trying to start server", err)
		}
	}()

	<-done

	fmt.Println("shutting down gracefully, press Ctrl+C again to force")
}

func initServer(port string, receiverUsecases usecase.ReceiverUseCases) {
	router := gin.Default()

	apiVersion1 := router.Group("api/v1")
	apiVersion1.POST("/receiver", graphqlHandler(receiverUsecases))
	apiVersion1.GET("/playground", playgroundHandler())

	router.Run(port)
}

func graphqlHandler(receiverUsecases usecase.ReceiverUseCases) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ReceiverUseCases: receiverUsecases}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/api/v1/receiver")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
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
