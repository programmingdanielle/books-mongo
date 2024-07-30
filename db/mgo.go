package db

import (
	"context"
	"fmt"

	"github.com/programmingdanielle/books-mongo/configs"
	"github.com/programmingdanielle/books-mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Collection = "BookDetails"

	Database = "Books"
)

var (
	// A placeholder for an error
	ErrNoDocumentID = fmt.Errorf("no _id given")
)

var BookCollection *mongo.Collection = configs.GetCollection(configs.DB, "BookDetails")

func InsertBook(ctx context.Context, newBook models.Book) (result *mongo.InsertOneResult, err error) {
	result, err = BookCollection.InsertOne(ctx, newBook)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetBook(ctx context.Context, id primitive.ObjectID) (models.Book, error) {
	var err error

	book := models.Book{}

	if id.String() == "" {
		return models.Book{}, ErrNoDocumentID
	}

	filter := primitive.M{"_id": id}

	if err = BookCollection.FindOne(ctx, filter).Decode(&book); err != nil {
		return book, err
	}

	return book, nil
}

func Find(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	results, err := BookCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return []models.Book{}, err
	}

	for results.Next(context.TODO()) {
		var singleBook models.Book
		err = results.Decode(&singleBook)
		if err != nil {
			return []models.Book{}, err
		}
		books = append(books, singleBook)
	}

	return books, nil
}

func Update(ctx context.Context, request models.Book) (mongo.UpdateResult, error) {
	// update document
	result := &mongo.UpdateResult{}
	if request.ID == nil {
		return *result, ErrNoDocumentID
	}

	filter := primitive.M{"_id": request.ID}
	update := bson.M{"$set": request}
	opts := options.Update().SetUpsert(false)
	result, err := BookCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return mongo.UpdateResult{}, err
	}
	if result.MatchedCount < 1 {
		return *result, mongo.ErrNoDocuments
	}
	return *result, err
}

func Delete(ctx context.Context, objId primitive.ObjectID) (mongo.DeleteResult, error) {
	var err error
	result := &mongo.DeleteResult{}
	filter := primitive.M{"_id": objId}
	opts := options.Delete()
	result, err = BookCollection.DeleteOne(ctx, filter, opts)
	if err != nil {
		return mongo.DeleteResult{}, err
	}
	if result.DeletedCount < 1 {
		return *result, mongo.ErrUnacknowledgedWrite
	}

	return *result, err
}

// func DeleteAll(ctx context.Context, objId primitive.ObjectID) (mongo.DeleteResult, error) {
// 	var err error
// 	result := &mongo.DeleteResult{}
// 	filter := primitive.M{"_id": objId}
// 	opts := options.Delete()
// 	result, err = BookCollection.DeleteMany(ctx, filter, opts)
// 	if err != nil {
// 		return mongo.DeleteResult{}, err
// 	}
// 	if result.DeletedCount < 1 {
// 		return *result, mongo.ErrUnacknowledgedWrite
// 	}

// 	return *result, err
// }
