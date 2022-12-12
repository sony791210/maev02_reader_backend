package platform

import (
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
	return c.JSON(util.Success(repositories.GetPlatFormIntroduce(types)))
}

func LastInfo(c *fiber.Ctx) error {
	return c.JSON(util.Success(repositories.GetPlatFormLastInfo()))
}

func List(c *fiber.Ctx) error {
	return c.JSON(util.Success(repositories.GetPlatFormList()))
}

func BookList(c *fiber.Ctx) error {
	return c.JSON(util.Success(repositories.GetPlatFormList()))
}
