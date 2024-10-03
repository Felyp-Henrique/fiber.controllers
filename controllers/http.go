package controllers

import "strings"

const (
	HttpMethodGet    HttpMethod = "Get"
	HttpMethodPost   HttpMethod = "Post"
	HttpMethodPut    HttpMethod = "Put"
	HttpMethodDelete HttpMethod = "Delete"
)

type HttpMethod string

func (h HttpMethod) IsValid() bool {
	return !strings.EqualFold(h.String(), "")
}

func (h HttpMethod) IsEquals(other HttpMethod) bool {
	return strings.EqualFold(h.String(), other.String())
}

func (h HttpMethod) String() string {
	return string(h)
}

func NewHttpMethodFromString(method string) HttpMethod {
	switch strings.ToLower(method) {
	case "get":
		return HttpMethodGet
	case "post":
		return HttpMethodPost
	case "put":
		return HttpMethodPut
	case "delete":
		return HttpMethodDelete
	}
	return ""
}
