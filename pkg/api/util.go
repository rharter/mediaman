package api

import "net/http"

func getPathParam(r *http.Request, name string) string {
	return r.URL.Query().Get(":" + name)
}
