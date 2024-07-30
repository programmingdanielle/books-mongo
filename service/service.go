package routes

import (
	// golang internal packages
	"encoding/json"
	"net/http"

	// local packages
	ctrl "github.com/programmingdanielle/books-mongo/controller"
	"github.com/programmingdanielle/books-mongo/helpers"
	"github.com/programmingdanielle/books-mongo/models"

	// imported from third parties
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const id = "id"

func Routes(router *mux.Router) {
	router.HandleFunc("/createBook", create()).Methods("POST")
	router.HandleFunc("/book/{id}", get()).Methods("GET")
	router.HandleFunc("/books", findBooks()).Methods("GET")
	router.HandleFunc("/book/update/{id}", update()).Methods("PUT")

	// need to troubleshoot delete endpoint -- throws EOF without even trying the code base
	// router.HandleFunc("/book/delete/{id}", delete()).Methods("DELETE")
}

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// retrieve the request body
		var book models.Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		// attempt to create a book
		objId, err := ctrl.Insert(ctx, book)
		if ctx.Err() != nil {
			return
		}

		if err != nil {
			helpers.RespondWithError(w, http.StatusUnprocessableEntity, err)
		}

		helpers.RespondWithJSON(w, http.StatusCreated, models.IDResponse{ID: objId})
	}
}

func get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := r.Context()

		id := mux.Vars(r)[id]

		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		book, err := ctrl.Get(ctx, objId)
		if ctx.Err() != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, ctx.Err())
			return
		}

		if err != nil {
			switch err {
			case mongo.ErrNoDocuments:
				helpers.RespondWithError(w, http.StatusNotFound, err)
			default:
				helpers.RespondWithError(w, http.StatusUnprocessableEntity, err)
			}
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, book)
	}
}

func findBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		ctx := r.Context()

		books, err := ctrl.Find(ctx)

		if err != nil {
			switch err {
			case mongo.ErrNoDocuments:
				helpers.RespondWithError(w, http.StatusNotFound, err)
			default:
				helpers.RespondWithError(w, http.StatusUnprocessableEntity, err)
			}
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, books)
	}
}

func update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		id := mux.Vars(r)[id]
		ctx := r.Context()

		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		var request models.Book
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		err = ctrl.Update(ctx, objId, request)
		if ctx.Err() != nil {
			return
		}

		if err != nil {
			switch err {
			case mongo.ErrNoDocuments:
				helpers.RespondWithError(w, http.StatusNotFound, err)
			default:
				helpers.RespondWithError(w, http.StatusUnprocessableEntity, err)
			}
		}

		helpers.RespondWithJSON(w, http.StatusOK, models.IDResponse{ID: objId})
	}
}

func delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		id := mux.Vars(r)[id]
		ctx := r.Context()

		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		var request models.Book
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		result, err := ctrl.Delete(ctx, objId)
		if ctx.Err() != nil {
			return
		}

		if err != nil {
			switch err {
			case mongo.ErrNoDocuments:
				helpers.RespondWithError(w, http.StatusNotFound, err)
			default:
				helpers.RespondWithError(w, http.StatusUnprocessableEntity, err)
			}
		}

		helpers.RespondWithJSON(w, http.StatusOK, models.IDResponse{ID: result.DeletedCount})
	}
}
