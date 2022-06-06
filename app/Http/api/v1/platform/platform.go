package platform

import (
	"fmt"
	"test/app/util"

	"test/app/repositories"

	"github.com/gofiber/fiber/v2"
)

//搜尋DB裏面相關的資訊
func Search(c *fiber.Ctx) error {
	query := c.Query("query")
	return c.JSON(util.Success(repositories.GetPlatFormData(query)))
}

func Introduce(c *fiber.Ctx) error {

	types := c.Params("types")
	fmt.Println(types)
	return c.JSON(util.Success(repositories.GetPlatFormIntroduce(types)))
}
