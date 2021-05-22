package db

import (
	"aiqap-back/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://yernar:uuSDFddDDD7!@cluster0.wa4dr.mongodb.net/aiqap?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
	}
	return client
}

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	client := ConnectDb()
	if client == nil {
		fmt.Println("client is null")
		return books, nil
	}
	collection := client.Database("aiqap").Collection("books")
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return books, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)
	}

	if err := cursor.Err(); err != nil {
		return books, err
	}

	return books, nil
}

func CreateBook(book models.Book) *mongo.InsertOneResult {

	client := ConnectDb()
	collection := client.Database("aiqap").Collection("books")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, book)
	return result
}

func GetBook(id primitive.ObjectID) (models.Book, error) {
	var book models.Book

	client := ConnectDb()
	collection := client.Database("aiqap").Collection("books")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)

	return book, err
}

func DeleteBook(id primitive.ObjectID) bool {
	client := ConnectDb()
	collection := client.Database("aiqap").Collection("books")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		fmt.Println("cannot delete error: ", err)
		return false
	}

	fmt.Println("Successfully deleted")
	return true
}
