package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PageData struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func OKWithPage(c *gin.Context, list interface{}, total int64, page, limit int) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data: PageData{
			List:  list,
			Total: total,
			Page:  page,
			Limit: limit,
		},
	})
}

func Fail(c *gin.Context, httpCode int, bizCode int, message string) {
	c.JSON(httpCode, Response{
		Code:    bizCode,
		Message: message,
		Data:    nil,
	})
}

func BadRequest(c *gin.Context, message string) {
	Fail(c, http.StatusBadRequest, 40000, message)
}

func Unauthorized(c *gin.Context, message string) {
	Fail(c, http.StatusUnauthorized, 40100, message)
}

func Forbidden(c *gin.Context, message string) {
	Fail(c, http.StatusForbidden, 40300, message)
}

func NotFound(c *gin.Context, message string) {
	Fail(c, http.StatusNotFound, 40400, message)
}

func ServerError(c *gin.Context, message string) {
	Fail(c, http.StatusInternalServerError, 50000, message)
}
