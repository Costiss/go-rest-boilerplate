package httputil

import "github.com/gin-gonic/gin"

type HTTPError struct {
	Err error `json:"message"`
}

func NewError(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, HTTPError{Err: err})
}
