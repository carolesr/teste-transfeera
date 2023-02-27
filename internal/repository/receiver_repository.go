package repository

import (
	"context"
	"errors"
	"time"

	"github.com/teste-transfeera/internal/entity"
	"github.com/teste-transfeera/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReceiverRepository interface {
	Create(receiver entity.Receiver) (*entity.Receiver, error)
	List(filter map[string]string) ([]entity.Receiver, error)
	FindById(id string) (*entity.Receiver, error)
	Update(id string, fields map[string]string) error
	Delete(ids []string) error
}

type receiverRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewReceiverRepository(collection *mongo.Collection, ctx context.Context) ReceiverRepository {
	return &receiverRepository{
		collection: collection,
		ctx:        ctx,
	}
}

func (r *receiverRepository) Create(receiver entity.Receiver) (*entity.Receiver, error) {
	model := model.Receiver{
		ID:         primitive.NewObjectID(),
		Identifier: receiver.Identifier,
		Name:       receiver.Name,
		Email:      receiver.Email,
		Pix: model.Pix{
			KeyType: string(receiver.Pix.KeyType),
			Key:     receiver.Pix.Key,
		},
		Status:    string(receiver.Status),
		CreatedAt: time.Now(),
	}

	_, err := r.collection.InsertOne(r.ctx, &model)
	if err != nil {
		return nil, err
	}

	entity := model.ToEntity()
	return &entity, nil
}

func (r *receiverRepository) List(filter map[string]string) ([]entity.Receiver, error) {
	bsonFilter := buildFilter(filter)
	bsonFilter["deleted_at"] = bson.M{"$exists": false}
	findOptions := options.Find()

	cursor, err := r.collection.Find(r.ctx, bsonFilter, findOptions)
	if err != nil {
		return nil, err
	}

	var receivers []entity.Receiver
	for cursor.Next(r.ctx) {
		var receiver model.Receiver
		err := cursor.Decode(&receiver)
		if err != nil {
			return nil, err
		}

		entity := receiver.ToEntity()
		receivers = append(receivers, entity)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(r.ctx)

	return receivers, nil
}

func (r *receiverRepository) FindById(id string) (*entity.Receiver, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	bsonFilter := bson.M{"_id": docID, "deleted_at": bson.M{"$exists": false}}

	result := r.collection.FindOne(r.ctx, bsonFilter)

	var receiver model.Receiver
	err = result.Decode(&receiver)
	if err != nil {
		return nil, err
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	entity := receiver.ToEntity()
	return &entity, nil
}

func (r *receiverRepository) Update(id string, fields map[string]string) error {
	docID, err := primitive.ObjectIDFromHex(id)
	bsonFilter := bson.M{"_id": docID}
	updater := bson.D{
		{"$set", buildUpdate(fields)},
	}

	result, err := r.collection.UpdateOne(r.ctx, bsonFilter, updater)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("record does not exist")
	}

	return nil
}

func (r *receiverRepository) Delete(ids []string) error {
	bsonList := bson.A{}
	for _, id := range ids {
		docID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		idFilter := bson.M{"_id": docID}
		bsonList = append(bsonList, idFilter)
	}

	bsonFilter := bson.M{"$or": bsonList}
	updater := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "deleted_at",
					Value: time.Now(),
				},
			},
		},
	}

	result, err := r.collection.UpdateMany(r.ctx, bsonFilter, updater)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("record does not exist")
	}

	return nil
}

func buildFilter(filter map[string]string) bson.M {
	bsonFilter := bson.M{}

	if filter["status"] != "" {
		bsonFilter["status"] = filter["status"]
	}
	if filter["name"] != "" {
		bsonFilter["name"] = filter["name"]
	}
	if filter["key_type"] != "" {
		bsonFilter["pix.key_type"] = filter["key_type"]
	}
	if filter["key"] != "" {
		bsonFilter["pix.key"] = filter["key"]
	}

	return bsonFilter
}

func buildUpdate(fields map[string]string) bson.D {
	var bsonUpdate bson.D

	if fields["identifier"] != "" {
		bsonUpdate = append(bsonUpdate, primitive.E{"identifier", fields["identifier"]})
	}
	if fields["name"] != "" {
		bsonUpdate = append(bsonUpdate, primitive.E{"name", fields["name"]})
	}
	if fields["email"] != "" {
		bsonUpdate = append(bsonUpdate, primitive.E{"email", fields["email"]})
	}
	if fields["key_type"] != "" {
		bsonUpdate = append(bsonUpdate, primitive.E{"pix.key_type", fields["key_type"]})
	}
	if fields["key"] != "" {
		bsonUpdate = append(bsonUpdate, primitive.E{"pix.key", fields["key"]})
	}

	bsonUpdate = append(bsonUpdate, primitive.E{"updated_at", time.Now()})

	return bsonUpdate
}
