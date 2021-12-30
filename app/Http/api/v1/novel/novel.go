package novel

import (
	"strconv"
	"test/app/util"

	"test/app/repositories"

	"github.com/gofiber/fiber/v2"
)

// ListNovel godoc
// @Summary Show a listNovel
// @Description get string by ID
// @Tags Novel
// @version 1.0
// @Accept  json
// @Produce  json
// @Success 200 {object} Novel
// @Router /api/v1/listNovel [get]
func ListNovel(c *fiber.Ctx) error {
	return c.JSON(util.Success(repositories.GetNovelList()))
}

// ListNovel godoc
// @Summary Show a ListNovelTitle
// @Description get string by ID
// @Tags Novel
// @version 1.0
// @Accept  json
// @Produce  json
// @Param id path int true "novel name id"
// @Success 200 {object} Novel
// @Router /api/v1/ListNovelTitle/{id} [get]
func ListNovelTitle(c *fiber.Ctx) error {
	sid := c.Params("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		return c.JSON("error")
	}
	return c.JSON(util.Success(repositories.GetNovelTitleList(id)))
}

// ListNovel godoc
// @Summary Show a NovelContext
// @Description get string by ID
// @Tags Novel
// @version 1.0
// @Accept  json
// @Produce  json
// @Param id path int true "novel ID"
// @Success 200 {object} Novel
// @Router /api/v1/NovelContext/{id}  [get]
func NovelContext(c *fiber.Ctx) error {
	sid := c.Params("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		return c.JSON("error")
	}
	return c.JSON(util.Success(repositories.GetNovelContent(id)))
}

type Novel struct {
	Id string
}

// ListNovel godoc
// @Summary Show a NovelContext
// @Description get string by ID
// @Tags Novel
// @version 1.0
// @Accept  json
// @Produce  json
// @Param query query string true "title name"
// @Success 200 {object} Novel
// @Router /api/v1/book/search  [get]
func BookSearch(c *fiber.Ctx) error {
	query := c.Query("query")
	return c.JSON(util.Success(repositories.GetAllData(query)))
}

// ListNovel godoc
// @Summary 書
// @Description get string by ID
// @Tags Novel
// @version 1.0
// @Accept  json
// @Produce  json
// @Param id path int true "novel ID"
// @Success 200 {object} Novel
// @Router /api/v1/book/{id} [get]
func BookInfo(c *fiber.Ctx) error {
	sid := c.Params("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		return c.JSON("error")
	}
	return c.JSON(util.Success(repositories.GetInfoData(id)))
}

func BookDetail(c *fiber.Ctx) error {
	sid := c.Params("id")
	spage := c.Params("page")
	id, err := strconv.Atoi(sid)
	if err != nil {
		return c.JSON("error")
	}
	page, err := strconv.Atoi(spage)
	if err != nil {
		return c.JSON("error")
	}
	return c.JSON(util.Success(repositories.GetDetailData(id, page)))
}
