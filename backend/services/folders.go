package services

import (
	"main/backend/models"
	"strconv"
)

type FoldersService struct {
	BaseService
}

func (f FoldersService) GetFolders(pagination models.Pagination) (folders []models.Folder, updatedPagination models.Pagination, err error) {
	var count int64 = 120
	updatedPagination = pagination
	updatedPagination.Calculate(count)
	folders = []models.Folder{}

	for i := 0; i < 10; i++ {
		folders = append(folders, models.Folder{Name: "Folder " + strconv.Itoa(i)})
	}

	return
}
