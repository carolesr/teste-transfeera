package graph_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/teste-transfeera/internal/entity"
	"github.com/teste-transfeera/internal/graph"
	"github.com/teste-transfeera/internal/usecase"
	"github.com/teste-transfeera/mocks"
)

type graphQLRequest struct {
	Query string `json:"query"`
}

func Test_Resolvers_CreateReceiver_Success(t *testing.T) {
	useCase := &mocks.ReceiverUseCases{}
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ReceiverUseCases: useCase}}))
	router := gin.Default()
	router.POST("/api/v1", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	t.Run("Resolve CreateReceiver successfully", func(t *testing.T) {
		// Arrange
		input := graph.NewReceiver{
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			PixKeyType: "CPF",
			PixKey:     "111.111.111-11",
		}
		mockInput := &usecase.CreateReceiverInput{
			Identifier: input.Identifier,
			Name:       input.Name,
			Email:      input.Email,
			PixKeyType: input.PixKeyType,
			PixKey:     input.PixKey,
		}
		id := uuid.New().String()
		mockOutput := &entity.Receiver{
			ID:         id,
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Status:     entity.Draft,
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
		}

		expectedResult := graph.Receiver{
			ID:         id,
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Identifier: "111.111.111-11",
			Pix: &graph.Pix{
				KeyType: "CPF",
				Key:     "111.111.111-11",
			},
		}
		var result struct {
			Data struct {
				CreateReceiver graph.Receiver `json:"createReceiver"`
			} `json:"data"`
		}
		result.Data.CreateReceiver = expectedResult
		expectedResultBytes, err := json.Marshal(result)

		useCase.On("Create", mockInput).Return(mockOutput, nil).Once()

		// Act
		query := `
			mutation {
				createReceiver(input: {
					name: "%s",
					email: "%s",
					identifier: "%s",
					pixKeyType: "%s",
					pixKey: "%s",
					}) {
					id
					name
					email
					identifier
					pix {
						keyType
						key
					}
					bank
					agency
					account
					status
				}
			}
		`
		query = fmt.Sprintf(query, input.Name, input.Email, input.Identifier, input.PixKeyType, input.PixKey)
		gqlMarshalled, err := json.Marshal(graphQLRequest{Query: query})

		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/api/v1", strings.NewReader(string(gqlMarshalled)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResultBytes, rr.Body.Bytes())
		assert.Equal(t, http.StatusOK, rr.Code)
		useCase.AssertExpectations(t)
	})
}

func Test_Resolvers_ListReceivers_Success(t *testing.T) {
	useCase := &mocks.ReceiverUseCases{}
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ReceiverUseCases: useCase}}))
	router := gin.Default()
	router.POST("/api/v1", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	t.Run("Resolve ListReceivers successfully", func(t *testing.T) {
		// Arrange
		filter := graph.BuildFilter(nil, nil, nil, nil)
		id1 := uuid.New().String()
		id2 := uuid.New().String()
		mockOutput := []entity.Receiver{
			{
				ID:         id1,
				Identifier: "111.111.111-11",
				Name:       "Receiver 1",
				Email:      "RECEIVER1@GMAIL.COM",
				Status:     entity.Draft,
				Pix: entity.Pix{
					KeyType: entity.CPF,
					Key:     "111.111.111-11",
				},
			},
			{
				ID:         id2,
				Identifier: "222.222.222-22",
				Name:       "Receiver 2",
				Email:      "RECEIVER2@GMAIL.COM",
				Status:     entity.Draft,
				Pix: entity.Pix{
					KeyType: entity.CPF,
					Key:     "222.222.222-22",
				},
			},
		}
		var b bool = false
		expectedResult := &graph.Receivers{
			Edges: []*graph.Edge{
				{
					Cursor: graph.EncodeBase64([]byte(id1)),
					Node: &graph.Receiver{
						ID:         id1,
						Identifier: "111.111.111-11",
						Name:       "Receiver 1",
						Email:      "RECEIVER1@GMAIL.COM",
						Pix: &graph.Pix{
							KeyType: "CPF",
							Key:     "111.111.111-11",
						},
					},
				},
				{
					Cursor: graph.EncodeBase64([]byte(id2)),
					Node: &graph.Receiver{
						ID:         id2,
						Identifier: "222.222.222-22",
						Name:       "Receiver 2",
						Email:      "RECEIVER2@GMAIL.COM",
						Pix: &graph.Pix{
							KeyType: "CPF",
							Key:     "222.222.222-22",
						},
					},
				},
			},
			PageInfo: &graph.PageInfo{
				StartCursor: graph.EncodeBase64([]byte(id1)),
				EndCursor:   graph.EncodeBase64([]byte(id2)),
				HasNextPage: &b,
			},
		}
		var result struct {
			Data struct {
				ListReceivers *graph.Receivers `json:"listReceivers"`
			} `json:"data"`
		}
		result.Data.ListReceivers = expectedResult
		expectedResultBytes, err := json.Marshal(result)

		useCase.On("List", filter).Return(mockOutput, nil).Once()

		// Act
		query := `
			query {
				listReceivers {
					edges {
						cursor
						node {
							id
							name
							email
							identifier
							pix {
								keyType
								key
							}
							bank
							agency
							account
							status
						}
					}
					pageInfo {
						startCursor
						endCursor
						hasNextPage
					}
				}
			}
		`
		gqlMarshalled, err := json.Marshal(graphQLRequest{Query: query})

		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/api/v1", strings.NewReader(string(gqlMarshalled)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, string(expectedResultBytes), string(rr.Body.Bytes()))
		assert.Equal(t, http.StatusOK, rr.Code)
		useCase.AssertExpectations(t)
	})
}
