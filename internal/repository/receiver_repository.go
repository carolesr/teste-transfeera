package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/teste-transfeera/internal/entity"
	"github.com/teste-transfeera/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReceiverRepository interface {
	Create(receiver entity.Receiver) (*entity.Receiver, error)
	List() ([]entity.Receiver, error)
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
	model := ToModel(receiver)
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	model.ID = uuid.New()

	_, err := r.collection.InsertOne(r.ctx, &model)
	if err != nil {
		return nil, err
	}

	entity := model.ToEntity()
	return &entity, nil
}

func (r *receiverRepository) List() ([]entity.Receiver, error) {
	filter := bson.D{{}}
	findOptions := options.Find()
	cursor, err := r.collection.Find(r.ctx, filter, findOptions)
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

func ToModel(entity entity.Receiver) model.Receiver {
	return model.Receiver{
		Identifier: entity.Identifier,
		Name:       entity.Name,
		Email:      entity.Email,
		Pix: model.Pix{
			KeyType: string(entity.Pix.KeyType),
			Key:     entity.Pix.Key,
		},
		Bank:    entity.Bank,
		Agency:  entity.Agency,
		Account: entity.Account,
		Status:  string(entity.Status),
	}
}
