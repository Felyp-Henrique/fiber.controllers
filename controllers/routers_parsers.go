package controllers

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

type RouterParser struct {
}

func (r RouterParser) CheckHandlerHttpMethod(controllerMethod reflect.Method, httpMethod HttpMethod) bool {
	pathMethod := r.GetPathMethod(controllerMethod)
	return pathMethod.IsValid() && r.GetPathMethod(controllerMethod).IsEquals(httpMethod)
}

func (r RouterParser) IsHandlerGetHttpMethod(controllerMethod reflect.Method) bool {
	return r.CheckHandlerHttpMethod(controllerMethod, HttpMethodGet)
}

func (r RouterParser) IsHandlerPostHttpMethod(controllerMethod reflect.Method) bool {
	return r.CheckHandlerHttpMethod(controllerMethod, HttpMethodPost)
}

func (r RouterParser) IsHandlerPutHttpMethod(controllerMethod reflect.Method) bool {
	return r.CheckHandlerHttpMethod(controllerMethod, HttpMethodPut)
}

func (r RouterParser) IsHandlerDeleteHttpMethod(controllerMethod reflect.Method) bool {
	return r.CheckHandlerHttpMethod(controllerMethod, HttpMethodDelete)
}

func (r RouterParser) GetParamsRegexp() *regexp.Regexp {
	return regexp.MustCompile(`By(\w+)`)
}

func (r RouterParser) GetPathRegexp() *regexp.Regexp {
	return regexp.MustCompile(`Path(\w+)`)
}

func (r RouterParser) GetMethodRegexp() *regexp.Regexp {
	return regexp.MustCompile(`(Get|Post|Put|Delete)`)
}

func (r RouterParser) GetPathParams(controllerMethod reflect.Method) []string {
	paramsRegexpResults := r.GetParamsRegexp().FindAllString(controllerMethod.Name, -1)
	if len(paramsRegexpResults) > 0 {
		paramsRegexpResultsWithoutPrefixBy := strings.ReplaceAll(paramsRegexpResults[0], "By", "")
		paramsRegexpResultsWithoutPrefixBy = strings.ToLower(paramsRegexpResultsWithoutPrefixBy)
		params := strings.Split(paramsRegexpResultsWithoutPrefixBy, "and")
		return params
	}
	return []string{}
}

func (r RouterParser) GetPathName(controllerMethod reflect.Method) string {
	controllerMethodNameWithoutParams := r.GetParamsRegexp().ReplaceAllString(controllerMethod.Name, "")
	pathRegexpResults := r.GetPathRegexp().FindAllString(controllerMethodNameWithoutParams, -1)
	if len(pathRegexpResults) > 0 {
		pathRegexpResultsWithoutPrefixPath := strings.ReplaceAll(pathRegexpResults[0], "Path", "")
		pathRegexpResultsWithoutPrefixPath = strings.ToLower(pathRegexpResultsWithoutPrefixPath)
		return fmt.Sprintf("/%s", pathRegexpResultsWithoutPrefixPath)
	}
	return ""
}

func (r RouterParser) GetPathMethod(controllerMethod reflect.Method) HttpMethod {
	controllerMethodNameWithJustMethod := r.GetParamsRegexp().ReplaceAllString(controllerMethod.Name, "")
	controllerMethodNameWithJustMethod = r.GetPathRegexp().ReplaceAllString(controllerMethodNameWithJustMethod, "")
	return NewHttpMethodFromString(controllerMethodNameWithJustMethod)
}

func (r RouterParser) GetPathComplete(controllerMethod reflect.Method) string {
	pathName := r.GetPathName(controllerMethod)
	pathParams := []string{}
	for _, param := range r.GetPathParams(controllerMethod) {
		pathParams = append(pathParams, fmt.Sprintf(":%s", param))
	}
	newPath, _ := url.JoinPath(pathName, pathParams...)
	return newPath
}

func NewRouterParser() *RouterParser {
	return &RouterParser{}
}
