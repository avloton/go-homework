package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_helloHandler(t *testing.T) {

	type wantStruct struct {
		code int
		response string
		contentType string
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		want wantStruct
	}{
		// TODO: Add test cases.
		{
			name: "Simple positive test",
			want: wantStruct{
				code: 200,
				response: `{"message": "hello, world!"}`,
				contentType: "application/json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/hello", nil)
			w := httptest.NewRecorder()
			helloHandler(w, request)
			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, tt.want.code, res.StatusCode)
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.JSONEq(t, tt.want.response, string(resBody))
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))

		})
	}
}
