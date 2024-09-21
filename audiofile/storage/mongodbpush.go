package storage

import (
	"audiofile/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"os"

	"github.com/joho/godotenv"
)

func (f FlatFile) PushToMongoDB(audio *models.Audio) error {

	err := godotenv.Load(".env")

  if err != nil {
    return err
  }
	mongoURL := os.Getenv("MONGO_DB_URL")
  databaseName := os.Getenv("MONGODB_DATABASE")
  collectionName := os.Getenv("MONGODB_COLLECTION")

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
  data, err := audio.JSON()

  // data is a string we have to convert i correct fomat to insert in mongodb
  
  fmt.Println("Data to be inserted", data)
    
  // Convert JSON string to BSON document
  var bsonData bson.M
  if err := bson.UnmarshalExtJSON([]byte(data), true, &bsonData); err != nil {
      fmt.Println("Error unmarshalling JSON to BSON", err)
      return err
  }
  // Remove the Path key from the BSON document
  delete(bsonData, "Path")
  // Insert the BSON document into the collection
  collection := client.Database(databaseName).Collection(collectionName)
  _, err = collection.InsertOne(context.TODO(), bsonData)
  if err != nil {
      return err
  }
  fmt.Println("Inserted a single document: ", bsonData)
  return nil
}
