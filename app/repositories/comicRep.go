package repositories

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"test/app/models"
)

func GetComicPageList(comicname string, page int) interface{} {
	var filePath []string

	var infoData models.Comic

	//轉換 url 亂碼轉換中文
	url2Name, err := url.QueryUnescape(comicname)
	if err != nil {
		fmt.Println(err)
	}

	err2 := DB.Table("comic").Where("comic_name_id = ? AND page = ?", url2Name, page-1).First(&infoData)

	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(infoData.Path)

	// 進入DB找路徑出來
	files, err3 := ioutil.ReadDir(infoData.Path)
	//
	if err3 != nil {
		fmt.Println(err3)
	}

	//for _, s := range comicname {
	//	fmt.Printf("%c\t%d\t%T\n", s, s, s)
	//}

	//dirs, err := ioutil.ReadDir("./public/" + url2Name)
	//
	//for _, dir := range dirs {
	//	fmt.Println(dir.Name())
	//}

	//if err3 != nil {
	//	fmt.Println(err3)
	//}
	//
	for _, file := range files {
		filePath = append(filePath, infoData.Path+"/"+file.Name())
	}
	//

	outputData := models.ComicListInfo{
		FilePath: filePath,
		Title:    infoData.Title,
		Page:     page,
	}

	return outputData

}

func GetComicInfoData(comicname string) interface{} {
	var infoData models.ComicInfo4

	//轉換 url 亂碼轉換中文
	url2Name, errurl2Name := url.QueryUnescape(comicname)
	if errurl2Name != nil {
		fmt.Println(errurl2Name)
	}

	err := DB.Table("platform_info").Select("platform_info.*").
		Where("platform_info.comic_name_id = ?", url2Name).First(&infoData)

	if err != nil {
		fmt.Println(err)
	}

	return infoData
}
