package controllers

import (
	"gin-boilerplate/internal/models"
	"gin-boilerplate/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(ctx *gin.Context) {
	var example []*models.Example
	repository.Get(&example)
	ctx.JSON(http.StatusOK, &example)
}
