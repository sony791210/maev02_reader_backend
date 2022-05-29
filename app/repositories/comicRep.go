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

	err2 := DB.Table("comic").Where("name = ? AND page = ?", url2Name, page).First(&infoData)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(infoData.Title)

	// 進入DB找路徑出來
	files, err3 := ioutil.ReadDir("./public/" + url2Name + "/" + infoData.Title)
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
		filePath = append(filePath, "/static/"+url2Name+"/"+infoData.Title+"/"+file.Name())
	}
	//
	var data interface{}

	data = GetComicInfoData(url2Name)

	outputData := models.ComicListInfo{
		FilePath: filePath,
		Title:    data.(models.ComicInfo4).Title,
		Page:     page,
	}

	return outputData

}

func GetComicInfoData(id string) interface{} {
	var infoData models.ComicInfo4

	err := DB.Table("novel_info").Select("novel_info.*").
		Where("novel_info.comic_name_id = ?", id).First(&infoData)

	if err != nil {
		fmt.Println(err)
	}

	return infoData
}
