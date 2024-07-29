package controllers

import (
	"context"

	"github.com/programmingbunny/books/db"
	"github.com/programmingbunny/books/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(ctx context.Context, request models.Book) (*mongo.InsertOneResult, error) {
	request.ID = nil
	return db.InsertBook(ctx, request)
}

func Get(ctx context.Context, objId primitive.ObjectID) (models.Book, error) {
	return db.GetBook(ctx, objId)
}

func Find(ctx context.Context) ([]models.Book, error) {
	var books []models.Book

	books, err := db.Find(ctx)
	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func Update(ctx context.Context, objId primitive.ObjectID, request models.Book) error {
	_, err := db.Update(ctx, request)
	return err
}

func Delete(ctx context.Context, objId primitive.ObjectID) (mongo.DeleteResult, error) {
	return db.Delete(ctx, objId)
}
