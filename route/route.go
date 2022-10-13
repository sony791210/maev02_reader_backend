package route

import (
	"github.com/gofiber/fiber/v2"
	"test/app/Http/api/v1/account"
	"test/app/Http/api/v1/comic"
	"test/app/Http/api/v1/novel"
	"test/app/Http/api/v1/platform"
)

func SetApiRoutes(app *fiber.App) {
	api := app.Group("/api")
	{
		v1 := api.Group("/v1", middleware1)
		{
			//拿取所有platform的資料
			v1.Get("/list/all", platform.List)

			v1.Get("/listNovel", novel.ListNovel)
			v1.Get("/ListNovelTitle/:id", novel.ListNovelTitle)
			v1.Get("/NovelContext/:id", novel.NovelContext)
			v1.Get("/accounts/:id", account.ShowAccount)

			//平台相關
			v1.Get("/book/search", novel.BookSearch)
			v1.Get("/platform/search", platform.Search)

			v1.Get("/platform/introduce/:types", platform.Introduce)
			v1.Get("/platform/lastinfo", platform.LastInfo)
			//小說相關
			v1.Get("/book/:id", novel.BookInfo)
			v1.Get("/book/:id/:page", novel.BookDetail)
			v1.Get("/bookList/:id", novel.BookList)
			//漫畫相關
			v1.Get("/comic/:id", comic.BookInfo)
			v1.Get("/comic/:comicname/:page", comic.PageList)

		}
	}
}

func middleware1(c *fiber.Ctx) error {
	//log.Println("test")
	// 执行该中间件之后的逻辑
	return c.Next()
}
