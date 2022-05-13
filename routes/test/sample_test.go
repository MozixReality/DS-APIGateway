package routes_test

import (
	"APIGateway/constant"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	tests := []test{
		{
			name: "SUCCESS",
			path: "/api/v1/sample",
			want: constant.Response{
				Code: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, tt.path, nil)
			router.ServeHTTP(w, r)
			assert.Equal(t, http.StatusOK, w.Code)

			var resp constant.Response
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, tt.want.Code, resp.Code)

			if tt.want.Code == http.StatusOK {
				assert.Equal(t, tt.want.Data, resp.Data)
			}
		})
	}
}
