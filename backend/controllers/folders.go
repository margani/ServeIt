package controllers

import (
	"main/backend/models"
	"main/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FoldersController struct {
}

var foldersService = services.FoldersService{}

func (controller *FoldersController) GetFolders(context *gin.Context) {
	pagination := models.Pagination{}.Get(context)

	folders, pagination, err := foldersService.GetFolders(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"records":    folders,
		"pagination": pagination,
	})
}
