package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	models "github.com/nikhilrana/Golang-Tutorials/24-MongodbConnection/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString string
var databaseName string
var collectionName string

var Collection *mongo.Collection

func setEnvironmentValues() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	connectionString = os.Getenv("CONNECTION_URL")
	databaseName = os.Getenv("DATABASE_NAME")
	collectionName = os.Getenv("COLLECTION_NAME")
}

func init() {
	setEnvironmentValues()

	//client-option
	clientOption := options.Client().ApplyURI(connectionString)

	// connection to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	Collection = (*mongo.Collection)(client.Database(databaseName).Collection(collectionName))

	fmt.Println("Collection instance is ready")

}

func insertMovie(movie models.Netflix) {
	inserted, err := Collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie is inserted in db with id: ", inserted.InsertedID)
}

func updateMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	// result tells you how many value modified.

	fmt.Println("Total modifications: ", result.MatchedCount)
}

func deleteMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total deletion occurs: ", result.DeletedCount)
}

func deleteAllMovies() int64 {
	filter := bson.D{{}}
	result, err := Collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total no of deletion occurs: ", result.DeletedCount)
	return result.DeletedCount
}

func getAllMovies() []primitive.M {
	cursor, err := Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	// cursor points to all the data values and we iterate it to extract them

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}

// Controllers

func GetAllMoviesService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func AddMovieService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var newMovie models.Netflix

	json.NewDecoder(r.Body).Decode(&newMovie)
	insertMovie(newMovie)
	json.NewEncoder(w).Encode(newMovie)

}

func MarkedMovieAsWatchedService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateMovie(params["id"])
	json.NewEncoder(w).Encode("Movie updated succesfully")

}

func DeleteMovieService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteMovie(params["id"])
	json.NewEncoder(w).Encode("Movie deleted successfully")
}

func DeleteAllMoviesService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)

}
