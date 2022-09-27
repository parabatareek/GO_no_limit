package httptest_1

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserViewHandler(t *testing.T) {
	// Структруа ожидаемых результатов
	type want struct {
		statusCode  int64
		contentType string
		user        User
	}

	// Структура тестов
	tests := []struct {
		name    string
		request string
		users   map[string]User
		want    want
	}{
		{"simple test #1",
			"/users?user_id=1",
			map[string]User{"id1": {"id1",
				"Misha",
				"Potapov"}},
			want{
				statusCode:  200,
				contentType: "application/json",
				user:        User{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.request, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(UserViewHandler(tt.users))
			h.ServeHTTP(w, request)
			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			userResult, err := io.ReadAll(result.Body)
			require.NoError(t, err)
			err = result.Body.Close()
			require.NoError(t, err)

			var user User
			err = json.Unmarshal(userResult, &user)
			require.NoError(t, err)

			assert.Equal(t, tt.want.user, user)
		})
	}
}
