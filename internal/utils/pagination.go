package utils

import (
	"basic-crud/internal/models"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	limit := 10
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}

	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

func GeneratePaginateFromResponse(pagination models.Pagination, total int64) models.HTTPResponsePaginate {
	var limit float64 = float64(pagination.Limit)
	var count float64 = float64(total)
	totalPages := math.Ceil((count / limit))
	return models.HTTPResponsePaginate{
		Page:        pagination.Page,
		PerPage:     pagination.Limit,
		CurrentPage: pagination.Page,
		TotalPage:   int(totalPages),
	}
}
