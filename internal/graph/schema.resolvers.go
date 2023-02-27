package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/teste-transfeera/internal/usecase"
	"github.com/teste-transfeera/pkg/shared"
)

// CreateReceiver is the resolver for the createReceiver field.
func (r *mutationResolver) CreateReceiver(ctx context.Context, input NewReceiver) (*Receiver, error) {
	usecaseInput := &usecase.CreateReceiverInput{
		Name:       input.Name,
		Email:      input.Email,
		Identifier: input.Identifier,
		PixKeyType: input.PixKeyType,
		PixKey:     input.PixKey,
	}

	result, err := r.ReceiverUseCases.Create(usecaseInput)
	if err != nil {
		return nil, err
	}

	return ToOutput(*result), nil
}

// DeleteReceivers is the resolver for the deleteReceivers field.
func (r *mutationResolver) DeleteReceivers(ctx context.Context, ids []string) (string, error) {
	usecaseInput := &usecase.DeleteReceiverInput{
		Ids: ids,
	}

	err := r.ReceiverUseCases.Delete(usecaseInput)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Deleted %s successfully", ids)
	return result, nil
}

// UpdateReceiver is the resolver for the updateReceiver field.
func (r *mutationResolver) UpdateReceiver(ctx context.Context, input UpdateReceiver) (string, error) {
	usecaseInput := &usecase.UpdateReceiverInput{
		Id:         input.ID,
		Name:       shared.GetValueStr(input.Name),
		Email:      shared.GetValueStr(input.Email),
		Identifier: shared.GetValueStr(input.Identifier),
		PixKeyType: shared.GetValueStr(input.PixKeyType),
		PixKey:     shared.GetValueStr(input.PixKey),
	}

	err := r.ReceiverUseCases.Update(usecaseInput)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Updated %s successfully", input.ID)
	return result, nil
}

// Receiver is the resolver for the receiver field.
func (r *queryResolver) Receiver(ctx context.Context, id string) (*Receiver, error) {
	usecaseInput := &usecase.ListReceiverByIdInput{
		Id: id,
	}

	result, err := r.ReceiverUseCases.ListById(usecaseInput)
	if err != nil {
		return nil, err
	}

	return ToOutput(*result), nil
}

// ListReceivers is the resolver for the listReceivers field.
func (r *queryResolver) ListReceivers(ctx context.Context, first *int, after *string, status *string, name *string, keyType *string, key *string) (*Receivers, error) {
	filter := BuildFilter(status, name, keyType, key)
	receivers, err := r.ReceiverUseCases.List(filter)
	if err != nil {
		return nil, err
	}

	if len(receivers) == 0 {
		return &Receivers{
			Edges:    []*Edge{},
			PageInfo: &PageInfo{},
		}, nil
	}

	totalPerPage := TOTAL_PER_PAGE
	if first != nil {
		totalPerPage = *first
	}

	isInCurrentPage := true
	var cursor string
	if after != nil {
		cursor, err = shared.DecodeBase64(*after)
		if cursor != "" {
			isInCurrentPage = false
		}
	}

	count := 0
	hasNextPage := false
	edges := make([]*Edge, totalPerPage)
	for i, receiver := range receivers {
		hasReachedTotalPerPage := count == totalPerPage

		if isInCurrentPage && !hasReachedTotalPerPage {
			edges[count] = &Edge{
				Cursor: shared.EncodeBase64([]byte(receiver.ID)),
				Node:   ToOutput(receiver),
			}
			count++
		}

		if receiver.ID == cursor {
			isInCurrentPage = true
		}

		if hasReachedTotalPerPage {
			hasNextPage = len(receivers) > i
			break
		}
	}

	pageInfo := PageInfo{
		StartCursor: shared.EncodeBase64([]byte(edges[0].Node.ID)),
		EndCursor:   shared.EncodeBase64([]byte(edges[count-1].Node.ID)),
		HasNextPage: &hasNextPage,
	}

	result := Receivers{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return &result, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
