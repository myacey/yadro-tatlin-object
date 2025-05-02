package handler

import (
	"log"
	"net/http"

	"yadro-tatlin-object/internal/models/dto/response"
	"yadro-tatlin-object/internal/pkg/web/apperror"

	"github.com/gin-gonic/gin"
)

const (
	HeaderRequestID  = "X-Request-Id"
	CtxKeyRetryAfter = "Retry-After"
)

type Handler struct {
	petSrv PetService
}

func NewHandler(petSrv PetService) *Handler {
	return &Handler{petSrv: petSrv}
}

func wrapCtxWithError(ctx *gin.Context, err error) {
	if httpError, ok := err.(apperror.HTTPError); ok {
		ctx.JSON(httpError.Code, response.Error{
			Code:      httpError.Code,
			Message:   httpError.Message,
			RequestID: ctx.GetHeader(HeaderRequestID),
		})

		if httpError.Code == http.StatusInternalServerError {
			log.Printf("internal error: %v | %v", httpError.Message, httpError.DebugError)
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, response.Error{
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
			RequestID: ctx.GetHeader(HeaderRequestID),
		})
	}
	ctx.Set(CtxKeyRetryAfter, 10)
}
