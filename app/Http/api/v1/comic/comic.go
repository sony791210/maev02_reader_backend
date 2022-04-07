package comic

import (
	"test/app/util"

	"test/app/repositories"

	"github.com/gofiber/fiber/v2"
)

// ListNovel godoc
// @Summary Show a comic page list
// @Description get string by ID
// @Tags Comic
// @version 1.0
// @Accept  json
// @Produce  json
// @Param comicname path string true "comicname"
// @Param page path string true "page"
// @Success 200 {object} Comic
// @Router /api/v1/comic/list/{comicname}/{page} [get]
func PageList(c *fiber.Ctx) error {
	comicname := c.Params("comicname")
	page := c.Params("page")

	return c.JSON(util.Success(repositories.GetComicPageList(comicname, page)))
}




func BookInfo(c *fiber.Ctx) error {
	comicId := c.Params("id")
	return c.JSON(util.Success(repositories.GetComicInfoData(comicId)))

}

type Comic struct {
	Id string
}
