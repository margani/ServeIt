package services

import (
	"main/backend/models"
)

type FoldersService struct {
	BaseService
}

func (f FoldersService) GetFolders(pagination models.Pagination) (folders []models.Folder, updatedPagination models.Pagination, err error) {
	var count int64 = 3
	updatedPagination = pagination
	updatedPagination.Calculate(count)
	folders = []models.Folder{}

	folders = append(folders, models.Folder{
		Name: "Music",
		Path: "C:\\Users\\dev\\Music",
	})

	folders = append(folders, models.Folder{
		Name: "Videos",
		Path: "C:\\Users\\dev\\Videos",
	})

	folders = append(folders, models.Folder{
		Name: "Pictures",
		Path: "C:\\Users\\dev\\OneDrive\\Pictures",
	})

	return
}
