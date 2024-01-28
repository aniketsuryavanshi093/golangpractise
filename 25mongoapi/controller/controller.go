package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://aniketsuryavanshi093:kMsaFYSHPe1MU1Bl@golangpractise.fy4qwfr.mongodb.net/"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with monogoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

// MONGODB helpers - file

// insert 1 record
// func insertOneMovie(movie model.Netflix) {
// 	inserted, err := collection.InsertOne(context.Background(), movie)

//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
//	}
func insertOneMovie(movie model.Netflix) (model.Netflix, error) {
	result, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		return model.Netflix{}, err
	}

	// Retrieve the inserted document by querying with the InsertedID
	var insertedMovie model.Netflix
	err = collection.FindOne(context.Background(), bson.M{"_id": result.InsertedID}).Decode(&insertedMovie)
	if err != nil {
		return model.Netflix{}, err
	}

	fmt.Println("Inserted 1 movie in db with id: ", result.InsertedID)

	return insertedMovie, nil
}

// update 1 record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MOvie got delete with delete count: ", deleteCount)
}

// delete all records from mongodb
func deleteAllMovie() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NUmber of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
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

func getOneMovie(movieId string) primitive.M {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	var movie primitive.M
	err := collection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	return movie
}

// Actual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}

func GetAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	allmovies := getOneMovie(params["id"])
	json.NewEncoder(w).Encode(allmovies)
}

// func CreateMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/x-www-form-urlencode")
// 	w.Header().Set("Allow-Control-Allow-Methods", "POST")

// 	var movie model.Netflix

// 	_ = json.NewDecoder(r.Body).Decode(&movie)
// 	insertOneMovie(movie)
// 	json.NewEncoder(w).Encode(movie)
// }

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie model.Netflix

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertedMovie, err := insertOneMovie(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the inserted movie as JSON response
	json.NewEncoder(w).Encode(insertedMovie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
