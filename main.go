package main

import (
	"context"
	"fmt"
	"log"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)
collection := helper.ConnectDB()
type User struct{
	Id bson.ObjectId `json: "id" bson:"_id"'`
	Name string    `json:"name" bson:"name"`
	Email string	`json:"email" bson:"email"`
	Password string	`json:"password" bson:"password"`
}

type Post struct{
    User
	Id bson.ObjectId `json: "id" bson:"_id"'`
	Caption string    `json:"name" bson:"name"`
	Image_URL string	`json:"image" bson:"image"`
	Created_at  *timestamp.Timestamp  `json:"created_at,omitempty" bson:"created_at,omitempty" `

func getUser(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.User

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&user)

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}
func main() {
	//Init Router
	r := mux.NewRouter()

  	// arrange our route
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
    r.HandleFunc("/api/posts/{id}", getPost).Methods("GET")
    r.HandleFunc("/api/posts/users/{id}", getPostofUser).Methods("GET")
	r.HandleFunc("/api/posts", createPost).Methods("POST")

  	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}
