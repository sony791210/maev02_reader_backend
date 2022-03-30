package route

import (
	"log"

	"test/app/Http/api/v1/account"
	"test/app/Http/api/v1/comic"
	"test/app/Http/api/v1/novel"

	"github.com/gofiber/fiber/v2"
)

func SetApiRoutes(app *fiber.App) {
	api := app.Group("/api")
	{
		v1 := api.Group("/v1", middleware1)
		{
			v1.Get("/listNovel", novel.ListNovel)
			v1.Get("/ListNovelTitle/:id", novel.ListNovelTitle)
			v1.Get("/NovelContext/:id", novel.NovelContext)
			v1.Get("/accounts/:id", account.ShowAccount)
			v1.Get("/book/search", novel.BookSearch)
			v1.Get("/book/:id", novel.BookInfo)
			v1.Get("/book/:id/:page", novel.BookDetail)
			v1.Get("/bookList/:id", novel.BookList)

			v1.Get("/comic/list/:comicname/:page", comic.PageList)
		}
	}
}

func middleware1(c *fiber.Ctx) error {
	log.Println("test")
	// 执行该中间件之后的逻辑
	return c.Next()
}
