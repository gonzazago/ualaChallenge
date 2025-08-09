package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContextKey string

const ParamsKey ContextKey = "params"

type Handler interface {
	Handle(respWriter http.ResponseWriter, req *http.Request)
}

// AdaptHandler convierte nuestro web.Handler en un gin.HandlerFunc.
func AdaptHandler(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := make(map[string]string)
		for _, p := range c.Params {
			params[p.Key] = p.Value
		}
		ctx := context.WithValue(c.Request.Context(), ParamsKey, params)
		h.Handle(c.Writer, c.Request.WithContext(ctx))
	}
}
