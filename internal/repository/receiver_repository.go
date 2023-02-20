package repository

import (
	"context"

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

	result, err := r.collection.InsertOne(r.ctx, &model)
	if err != nil {
		return nil, err
	}

	receiver.ID = result.InsertedID.(string)
	return &receiver, nil
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
		var receiver entity.Receiver
		err := cursor.Decode(&receiver)
		if err != nil {
			return nil, err
		}

		receivers = append(receivers, receiver)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(r.ctx)

	return receivers, nil
}

func ToModel(entity entity.Receiver) model.Receiver {
	return model.Receiver{
		ID:         entity.ID,
		Identifier: entity.Identifier,
		Name:       entity.Name,
		Email:      entity.Email,
		PixKeyType: entity.PixKeyType,
		PixKey:     entity.PixKey,
		Bank:       entity.Bank,
		Agency:     entity.Agency,
		Account:    entity.Account,
		Status:     entity.Status,
	}
}
