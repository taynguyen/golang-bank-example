package handlers

import (
	"errors"
	"gin-boilerplate/internal/models"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit  int   `form:"limit" json:"limit"`
	Offset int   `form:"offset" json:"offset"`
	Total  int64 `form:"total" json:"total"`
}

func GetPaginationFromQuery(c *gin.Context, defaultLimit int) (models.Pagination, error) {
	var p Pagination
	err := c.BindQuery(&p)
	if err != nil {
		return models.Pagination{Limit: defaultLimit}, errors.New("invalid pagination query params")
	}

	m := models.Pagination{
		Limit:  p.Limit,
		Offset: p.Offset,
	}
	if p.Limit == 0 {
		m.Limit = defaultLimit
	}
	return m, nil
}
