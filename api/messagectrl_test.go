package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hdlopez/clean-architecture-golang/apierror"
	"github.com/hdlopez/clean-architecture-golang/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockService is a mock implementation of message.Service
type MockService struct {
	mock.Mock
}

func (m *MockService) Get(ctx context.Context, id string) (*message.Message, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*message.Message), args.Error(1)
}

func TestMessageCtrl_Get(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		id             string
		mockReturnMsg  *message.Message
		mockReturnErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Success",
			id:             "123",
			mockReturnMsg:  &message.Message{Text: "Hello"},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"text":"Hello"}`,
		},
		{
			name:           "API Error",
			id:             "400",
			mockReturnMsg:  nil,
			mockReturnErr:  apierror.New(http.StatusBadRequest, "Bad Request"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"status":400,"message":"Bad Request"}`,
		},
		{
			name:           "Internal Server Error",
			id:             "500",
			mockReturnMsg:  nil,
			mockReturnErr:  errors.New("unexpected error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status":500,"message":"unexpected error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSrv := new(MockService)
			mockSrv.On("Get", mock.Anything, tt.id).Return(tt.mockReturnMsg, tt.mockReturnErr)

			ctrl := NewMessageController(mockSrv)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = gin.Params{{Key: "id", Value: tt.id}}
			
			// We need to pass the gin context as Ctx
			ctrl.Get(c)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
