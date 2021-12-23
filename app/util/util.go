package util

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/**
成功返回，你可以建其它返回，如失敗、錯誤
Successful return, you can create other returns, such as failure and error
*/
func Success(data interface{}) fiber.Map {

	return fiber.Map{
		"code":    http.StatusOK,
		"data":    data,
		"message": "請求成功",
	}

}
