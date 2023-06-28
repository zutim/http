package http

import "github.com/gin-gonic/gin"

type Route struct {
	Method  string
	URL     string
	Handler gin.HandlerFunc
}

type RouteGroup struct {
	Prefix   string
	Routes   []Route
	Handlers []gin.HandlerFunc
}

func RegisterRouteGroup(r *gin.Engine, group RouteGroup) {
	groupRouter := r.Group(group.Prefix)
	for _, handler := range group.Handlers {
		groupRouter.Use(handler)
	}
	for _, route := range group.Routes {
		groupRouter.Handle(route.Method, route.URL, route.Handler)
	}
}
