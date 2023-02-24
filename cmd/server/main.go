package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
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

	ctx := context.Background() //context.WithTimeout(context.Background(), 5*time.Second)

	collection := initDB(ctx)
	receiverRepository := repository.NewReceiverRepository(collection, ctx)

	receiverUsecases := usecase.NewReceiverUseCases(receiverRepository)

	initServer(port, receiverUsecases)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initServer(port string, receiverUsecases usecase.ReceiverUseCases) {
	router := gin.Default()

	apiVersion1 := router.Group("api/v1")
	apiVersion1.POST("/receiver", graphqlHandler(receiverUsecases))

	router.Run(port)
}

func graphqlHandler(receiverUsecases usecase.ReceiverUseCases) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ReceiverUseCases: receiverUsecases}}))

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
