package repositories

import (
	"fmt"
	"io/ioutil"
	"test/app/models"
)



func GetComicPageList(comicname string, page string) interface{} {
	var filePath []string
	

	files, err := ioutil.ReadDir("./public/" + comicname + "/" + page)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		filePath = append(filePath, "/static/"+comicname+"/"+page+"/"+file.Name())
	}

	// var infoData interface{}
	// infoData =GetComicInfoData(comicname)

	// var outputData := models.ComicListInfo{
	// 	FilePath:filePath,
	// 	Title:"123",
	// 	page:page
	// }
	



	return filePath
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
