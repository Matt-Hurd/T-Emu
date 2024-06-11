package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Err    int         `json:"err"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func JSONResponse(c *gin.Context, err int, errmsg string, data interface{}) {
	c.JSON(err, Response{
		Err:    err,
		ErrMsg: errmsg,
		Data:   data,
	})
}
