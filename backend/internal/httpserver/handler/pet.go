package handler

import (
	"context"
	"log"
	"net/http"

	"yadro-tatlin-object/internal/models/dto/request"
	"yadro-tatlin-object/internal/models/entity"
	"yadro-tatlin-object/internal/pkg/web/apperror"

	"github.com/gin-gonic/gin"
)

type PetService interface {
	GetPet(ctx context.Context) (*entity.Pet, error)
	UploadPet(ctx context.Context, req *request.UploadPet) error
}

func (h Handler) GetPet(ctx *gin.Context) {
	log.SetPrefix("http-server.handler.GetPet")

	pet, err := h.petSrv.GetPet(ctx)
	if err != nil {
		wrapCtxWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, pet.ToResponse())
}

func (h Handler) UploadPet(ctx *gin.Context) {
	log.SetPrefix("http-server.handler.UploadPet")

	var req request.UploadPet
	if err := ctx.ShouldBindJSON(&req); err != nil {
		wrapCtxWithError(ctx, apperror.NewBadReq("invalid request"))
		return
	}

	err := h.petSrv.UploadPet(ctx, &req)
	if err != nil {
		wrapCtxWithError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
