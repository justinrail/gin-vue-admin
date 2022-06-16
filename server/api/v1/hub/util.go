package hub

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func queryInt(c *gin.Context, param string) int {
	res, err := strconv.Atoi(c.Query(param))
	if err != nil {
		return 0
	}

	return res
}
