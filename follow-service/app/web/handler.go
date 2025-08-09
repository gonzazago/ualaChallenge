package web

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type contextKey string

const ParamsKey contextKey = "params"

type Handler interface {
	Handle(respWriter http.ResponseWriter, req *http.Request)
}

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
