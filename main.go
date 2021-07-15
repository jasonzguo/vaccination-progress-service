package main

import (
	"context"
	"log"
	"net/http"

	controller "github.com/jasonzguo/vaccination-progress-service/controllers"
	middleware "github.com/jasonzguo/vaccination-progress-service/middlewares"
	repo "github.com/jasonzguo/vaccination-progress-service/repos"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func initializeMongoClient() (*mongo.Client, context.Context) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}

func initializeLogger() (*zap.Logger, func()) {
	logger := zap.NewExample()
	undo := zap.ReplaceGlobals(logger)
	return logger, undo
}

func initializeRepo(client *mongo.Client) {
	database := client.Database("vaccination")
	progressionCollection := database.Collection("progression")
	repo.GetProgressionRepo().SetCollection(progressionCollection)
}

func main() {
	logger, undo := initializeLogger()
	defer logger.Sync()
	defer undo()

	client, ctx := initializeMongoClient()
	defer client.Disconnect(ctx)

	initializeRepo(client)

	ms := middleware.NewStack()
	ms.Use(middleware.Log)
	ms.Use(middleware.Authenticate)

	router := httprouter.New()
	router.GET("/", ms.Wrap(controller.GetProgressionController().FindAll))
	log.Fatal(http.ListenAndServe(":8080", router))
}
