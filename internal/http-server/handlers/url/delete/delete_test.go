package delete_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/http-server/handlers/url/delete"
	"url-shortener/internal/http-server/handlers/url/delete/mocks"
	"url-shortener/internal/logger/handlers/slogdiscard"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestDeleteHandler(t *testing.T) {
	cases := []struct {
		name      string
		alias     string
		url       string
		respError string
		mockError error
	}{
		{
			name:  "Success",
			alias: "test_alias",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			urlDeleterMock := mocks.NewURLDeleter(t)

			if tc.respError == "" || tc.mockError != nil {
				urlDeleterMock.On("DeleteURL", tc.alias).
					Return(tc.mockError).
					Once()
			}

			r := chi.NewRouter()
			r.Delete("/url/{alias}", delete.New(slogdiscard.NewDiscardLogger(), urlDeleterMock))
			ts := httptest.NewServer(r)
			defer ts.Close()

			url := ts.URL + "/url/" + tc.alias
			req, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			client := ts.Client()
			resp, err := client.Do(req)
			require.NoError(t, err)

			require.Equal(t, resp.StatusCode, http.StatusOK)
		})
	}
}
