package api

import (
	"io"
	"net/http"
	"strconv"
)

func getPathParam(r *http.Request, name string) string {
	return r.URL.Query().Get(":" + name)
}

func getIdFromRequest(r *http.Request) (int64, error) {
	idstr := r.URL.Query().Get(":id")
	return strconv.ParseInt(idstr, 10, 64)
}

// Put executes a PUT request using the default http client
func Put(url string, bodyType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	return http.DefaultClient.Do(req)
}

// Delete executes a DELETE request using the default http client
func Delete(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}
