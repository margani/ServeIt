package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Current    uint   `json:"current"`
	Size       uint   `json:"size"`
	TotalItems uint64 `json:"totalItems"`
}

func (Pagination) Get(context *gin.Context) (pagination Pagination) {
	current, err := strconv.ParseUint(context.Query("current"), 10, 32)
	if err != nil || current < 1 {
		current = 1
	}

	size, err := strconv.ParseUint(context.Query("size"), 10, 32)
	if err != nil || size < 10 || size > 100 {
		size = 10
	}

	pagination = Pagination{
		Current: uint(current),
		Size:    uint(size),
	}

	return
}

func (p *Pagination) Calculate(count int64) {
	p.TotalItems = uint64(count)
}
