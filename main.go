package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"Appointy-Instagram/handlers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func connectToDb() *mongo.Client {
	mongoClient, err := mongo.Connect(context.Background(), &options.ClientOptions{
		Auth: &options.Credential{
			Username: "mongoadmin",
			Password: "secret",
		},
	})
	if err != nil {
		log.Fatalf("Unable to connect to db\n[Error]: %v", err)
	}
    mongoClient.Database("Insta").CreateCollection(context.Background(), "Users")
	mongoClient.Database("Insta").CreateCollection(context.Background(), "Posts")

	return mongoClient
}
func main() {
	
	mongoClient := connectToDb()
    userCollection := mongoClient.Database("Insta").Collection("Users")
	postCollection := mongoClient.Database("Insta").Collection("Posts")
    userHandler := handlers.NewUserHandler(userCollection)
	postHandler := handlers.NewPostHandler(postCollection)
	postUserHandler := handlers.NewPostUserHandler(postCollection)
    http.Handle("/users/", userHandler)
	http.Handle("/posts/", postHandler)
	http.Handle("/posts/users/", postUserHandler)
    fmt.Println("Starting the server at localhost:8080 ...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}