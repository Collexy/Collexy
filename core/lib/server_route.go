package lib

import "net/http"

type ServerRoute struct {
	Path        string
	HandlerFunc http.HandlerFunc
	Methods     []string
}