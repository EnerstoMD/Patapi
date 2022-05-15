package utils

import (
	"lupus/patapi/pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationFromRequest(c *gin.Context) model.Pagination {
	limit := 1000
	page := 1
	query := c.Request.URL.Query()
	for key, val := range query {
		queryValue := val[len(val)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		}
	}

	return model.Pagination{
		Limit: limit,
		Page:  page,
	}
}
