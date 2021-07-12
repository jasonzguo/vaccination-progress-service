package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	controller "github.com/jasonzguo/vaccination-progress-service/controllers"
	repo "github.com/jasonzguo/vaccination-progress-service/repos"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initializeMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func initializeRepo(client *mongo.Client) {
	database := client.Database("vaccination")
	progressionCollection := database.Collection("progression")
	repo.GetProgressionRepo().SetCollection(progressionCollection)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queries := r.URL.Query()

	documents, err := controller.GetProgressionController().GetAll(r.Context(), queries.Get("lastId"))

	if err != nil {
		log.Fatal(fmt.Errorf("[Index] error in calling ProgressionController.GetAll  %v", err))
	}

	documentsJson, err := json.Marshal(documents)
	if err != nil {
		log.Fatal(fmt.Errorf("[Index] error in calling json.Marshal  %v", err))
	}

	fmt.Fprint(w, string(documentsJson))
}

func main() {
	client := initializeMongoClient()
	initializeRepo(client)

	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
