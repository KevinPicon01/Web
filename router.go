package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func (r Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	handler, methodExist, ok := r.FindRoute(request.URL.Path, request.Method)
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExist {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(writer, request)
}
func (r *Router) FindRoute(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, ok := r.rules[path][method]
	return handler, ok, exist
}
func (r *Router) AddRoute(pattern string, handler http.HandlerFunc) {

}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}
