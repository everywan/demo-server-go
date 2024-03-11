package rest

import "github.com/gin-gonic/gin"

type Router interface {
	gin.IRouter
}

func NewRouter() Router {
	engine := gin.New()
	engine.Use()
	return engine
}
