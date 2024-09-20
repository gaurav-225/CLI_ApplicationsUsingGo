package storage

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    
	"os"
	"github.com/joho/godotenv"
)

func PushToMongoDB() error {

	err := godotenv.Load(".env")

  if err != nil {
    return err
  }
	mongoURL := os.Getenv("MONGO_DB_URL")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  opts := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverAPI)
  // Create a new client and connect to the server
  client, err := mongo.Connect(context.TODO(), opts)
  if err != nil {
    return err
  }
  defer func() {
    if err = client.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()
  // Send a ping to confirm a successful connection
  if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
    panic(err)
  }
  fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
  return nil
}
