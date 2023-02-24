package graph_test

import (
	"encoding/json"
	"errors"
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
	router.POST("/api/v1/receiver", func(c *gin.Context) {
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
		req, err := http.NewRequest(http.MethodPost, "/api/v1/receiver", strings.NewReader(string(gqlMarshalled)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResultBytes, rr.Body.Bytes())
		assert.Equal(t, http.StatusOK, rr.Code)
		useCase.AssertExpectations(t)
	})
}

func Test_Resolvers_CreateReceiver_Error(t *testing.T) {
	useCase := &mocks.ReceiverUseCases{}
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ReceiverUseCases: useCase}}))
	router := gin.Default()
	router.POST("/api/v1/receiver", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	t.Run("Resolve CreateReceiver with error from usecase", func(t *testing.T) {
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
		expectedError := `{"errors":[{"message":"error","path":["createReceiver"]}],"data":{"createReceiver":null}}`

		useCase.On("Create", mockInput).Return(nil, errors.New("error")).Once()

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
		req, err := http.NewRequest(http.MethodPost, "/api/v1/receiver", strings.NewReader(string(gqlMarshalled)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, []byte(expectedError), rr.Body.Bytes())
		useCase.AssertExpectations(t)
	})
}

func Test_Resolvers_ListReceivers_Success(t *testing.T) {
	useCase := &mocks.ReceiverUseCases{}
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ReceiverUseCases: useCase}}))
	router := gin.Default()
	router.POST("/api/v1/receiver", func(c *gin.Context) {
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
		req, err := http.NewRequest(http.MethodPost, "/api/v1/receiver", strings.NewReader(string(gqlMarshalled)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, string(expectedResultBytes), string(rr.Body.Bytes()))
		assert.Equal(t, http.StatusOK, rr.Code)
		useCase.AssertExpectations(t)
	})

	t.Run("Resolve ListReceivers with 0 receivers", func(t *testing.T) {
		// Arrange
		filter := graph.BuildFilter(nil, nil, nil, nil)
		mockOutput := []entity.Receiver{}
		expectedResult := &graph.Receivers{
			Edges:    []*graph.Edge{},
			PageInfo: &graph.PageInfo{},
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
		req, err := http.NewRequest(http.MethodPost, "/api/v1/receiver", strings.NewReader(string(gqlMarshalled)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, string(expectedResultBytes), string(rr.Body.Bytes()))
		assert.Equal(t, http.StatusOK, rr.Code)
		useCase.AssertExpectations(t)
	})

	t.Run("Resolve ListReceivers with first 3 receivers", func(t *testing.T) {
		// Arrange
		filter := graph.BuildFilter(nil, nil, nil, nil)
		id1 := uuid.New().String()
		id2 := uuid.New().String()
		id3 := uuid.New().String()
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
			{
				ID:         id3,
				Identifier: "333.333.333-33",
				Name:       "Receiver 3",
				Email:      "RECEIVER3@GMAIL.COM",
				Status:     entity.Draft,
				Pix: entity.Pix{
					KeyType: entity.CPF,
					Key:     "333.333.333-33",
				},
			},
			{
				ID:         uuid.New().String(),
				Identifier: "444.444.444-4",
				Name:       "Receiver 4",
				Email:      "RECEIVER4@GMAIL.COM",
				Status:     entity.Draft,
				Pix: entity.Pix{
					KeyType: entity.CPF,
					Key:     "444.444.444-44",
				},
			},
			{
				ID:         uuid.New().String(),
				Identifier: "555.555.555-55",
				Name:       "Receiver 5",
				Email:      "RECEIVER5@GMAIL.COM",
				Status:     entity.Draft,
				Pix: entity.Pix{
					KeyType: entity.CPF,
					Key:     "555.555.555-55",
				},
			},
		}
		var b bool = true
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
				{
					Cursor: graph.EncodeBase64([]byte(id3)),
					Node: &graph.Receiver{
						ID:         id3,
						Identifier: "333.333.333-33",
						Name:       "Receiver 3",
						Email:      "RECEIVER3@GMAIL.COM",
						Pix: &graph.Pix{
							KeyType: "CPF",
							Key:     "333.333.333-33",
						},
					},
				},
			},
			PageInfo: &graph.PageInfo{
				StartCursor: graph.EncodeBase64([]byte(id1)),
				EndCursor:   graph.EncodeBase64([]byte(id3)),
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
				listReceivers(first: 3) {
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
		req, err := http.NewRequest(http.MethodPost, "/api/v1/receiver", strings.NewReader(string(gqlMarshalled)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, string(expectedResultBytes), string(rr.Body.Bytes()))
		assert.Equal(t, http.StatusOK, rr.Code)
		useCase.AssertExpectations(t)
	})

	t.Run("Resolve ListReceivers with next 3 receivers", func(t *testing.T) {
		// Arrange
		filter := graph.BuildFilter(nil, nil, nil, nil)
		id3 := uuid.New().String()
		id4 := uuid.New().String()
		id5 := uuid.New().String()
		mockOutput := []entity.Receiver{
			{
				ID:         uuid.New().String(),
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
				ID:         uuid.New().String(),
				Identifier: "222.222.222-22",
				Name:       "Receiver 2",
				Email:      "RECEIVER2@GMAIL.COM",
				Status:     entity.Draft,
				Pix: entity.Pix{
					KeyType: entity.CPF,
					Key:     "222.222.222-22",
				},
			},
			{
				ID:         id3,
				Identifier: "333.333.333-33",
				Name:       "Receiver 3",
				Email:      "RECEIVER3@GMAIL.COM",
				Status:     entity.Draft,
				Pix: entity.Pix{
					KeyType: entity.CPF,
					Key:     "333.333.333-33",
				},
			},
			{
				ID:         id4,
				Identifier: "444.444.444-44",
				Name:       "Receiver 4",
				Email:      "RECEIVER4@GMAIL.COM",
				Status:     entity.Draft,
				Pix: entity.Pix{
					KeyType: entity.CPF,
					Key:     "444.444.444-44",
				},
			},
			{
				ID:         id5,
				Identifier: "555.555.555-55",
				Name:       "Receiver 5",
				Email:      "RECEIVER5@GMAIL.COM",
				Status:     entity.Draft,
				Pix: entity.Pix{
					KeyType: entity.CPF,
					Key:     "555.555.555-55",
				},
			},
		}
		var b bool = false
		expectedResult := &graph.Receivers{
			Edges: []*graph.Edge{
				{
					Cursor: graph.EncodeBase64([]byte(id4)),
					Node: &graph.Receiver{
						ID:         id4,
						Identifier: "444.444.444-44",
						Name:       "Receiver 4",
						Email:      "RECEIVER4@GMAIL.COM",
						Pix: &graph.Pix{
							KeyType: "CPF",
							Key:     "444.444.444-44",
						},
					},
				},
				{
					Cursor: graph.EncodeBase64([]byte(id5)),
					Node: &graph.Receiver{
						ID:         id5,
						Identifier: "555.555.555-55",
						Name:       "Receiver 5",
						Email:      "RECEIVER5@GMAIL.COM",
						Pix: &graph.Pix{
							KeyType: "CPF",
							Key:     "555.555.555-55",
						},
					},
				},
			},
			PageInfo: &graph.PageInfo{
				StartCursor: graph.EncodeBase64([]byte(id4)),
				EndCursor:   graph.EncodeBase64([]byte(id5)),
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
				listReceivers(first: 3, after: "%s") {
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
		query = fmt.Sprintf(query, graph.EncodeBase64([]byte(id3)))
		gqlMarshalled, err := json.Marshal(graphQLRequest{Query: query})

		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/api/v1/receiver", strings.NewReader(string(gqlMarshalled)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, string(expectedResultBytes), string(rr.Body.Bytes()))
		assert.Equal(t, http.StatusOK, rr.Code)
		useCase.AssertExpectations(t)
	})
}

func Test_Resolvers_ListReceivers_Error(t *testing.T) {
	useCase := &mocks.ReceiverUseCases{}
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ReceiverUseCases: useCase}}))
	router := gin.Default()
	router.POST("/api/v1/receiver", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	t.Run("Resolve ListReceivers with error from usecase", func(t *testing.T) {
		// Arrange
		filter := graph.BuildFilter(nil, nil, nil, nil)
		expectedError := `{"errors":[{"message":"error","path":["listReceivers"]}],"data":{"listReceivers":null}}`

		useCase.On("List", filter).Return(nil, errors.New("error")).Once()

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
		req, err := http.NewRequest(http.MethodPost, "/api/v1/receiver", strings.NewReader(string(gqlMarshalled)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, []byte(expectedError), rr.Body.Bytes())
		useCase.AssertExpectations(t)
	})
}
