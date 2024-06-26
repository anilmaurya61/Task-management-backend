package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
    "os"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
  )
  
  var (
    DBUser string
    DBPass string
  ) 

  func ConnectionDb() {

	err := godotenv.Load("../../.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	DBUser = os.Getenv("DB_USER")
    DBPass = os.Getenv("DB_PASS")

	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.obgtu5m.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0",
	DBUser, DBPass)
	
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
  
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
	  panic(err)
	}
  
	defer func() {
	  if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	  }
	}()
  
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
	  panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
  }
  