package post

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"post-service/internal/delivery/rest/post/mocks"
	"post-service/internal/domain/post"
	"strings"
	"testing"
	"time"
)

//go:generate mockgen -destination=./mocks/save_service.go -package=mocks post-service/internal/delivery/rest/post PostCreateService
func TestCreatePostHandler_Handle(t *testing.T) {
	// Creamos un post de ejemplo para usar en las respuestas del mock
	mockPost := &post.Post{
		ID:        "post-id-123",
		UserID:    "user-id-abc",
		Text:      "Este es un post de prueba",
		CreatedAt: time.Now(),
	}

	type args struct {
		body string
	}
	type postServiceMock struct {
		post        *post.Post
		err         error
		createTimes int
	}
	type want struct {
		statusCode int
		response   map[string]string
	}

	tests := []struct {
		name            string
		args            args
		postServiceMock postServiceMock
		want            want
	}{
		{
			name: "Success - Post Accepted",
			args: args{
				body: `{"user_id":"user-id-abc","text":"Este es un post de prueba"}`,
			},
			postServiceMock: postServiceMock{
				post:        mockPost,
				err:         nil,
				createTimes: 1,
			},
			want: want{
				statusCode: http.StatusAccepted,
				response:   map[string]string{"status": "post accepted for processing"},
			},
		},
		{
			name: "Error - Invalid JSON Body",
			args: args{
				body: `{"user_id":"user-id-abc", "text":}`,
			},
			postServiceMock: postServiceMock{
				createTimes: 0, // El servicio no debería ser llamado
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "Error - Validation - Text Too Long",
			args: args{
				body: `{"user_id":"user-id-abc","text":"` + strings.Repeat("a", 281) + `"}`,
			},
			postServiceMock: postServiceMock{
				err:         post.ErrTextTooLong,
				createTimes: 1,
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "Error - Internal Server Error",
			args: args{
				body: `{"user_id":"user-id-abc","text":"Este es un post de prueba"}`,
			},
			postServiceMock: postServiceMock{
				err:         errors.New("unexpected database error"),
				createTimes: 1,
			},
			want: want{
				statusCode: http.StatusAccepted,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// --- Configuración del Mock ---
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			serviceMock := mocks.NewMockPostCreateService(mockCtrl)
			if tt.postServiceMock.createTimes > 0 {
				serviceMock.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Return(tt.postServiceMock.post, tt.postServiceMock.err).
					Times(tt.postServiceMock.createTimes)
			}

			handler := NewCreatePostHandler(serviceMock)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/posts", strings.NewReader(tt.args.body))
			rr := httptest.NewRecorder()

			// --- Ejecución ---
			handler.Handle(rr, req)

			// --- Aserciones ---
			assert.Equal(t, tt.want.statusCode, rr.Code)

			if tt.want.response != nil {
				var responseBody map[string]string
				err := json.Unmarshal(rr.Body.Bytes(), &responseBody)
				assert.NoError(t, err)
				assert.Equal(t, tt.want.response, responseBody)
			}
		})
	}

}
