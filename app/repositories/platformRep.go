package repositories

import (
	"fmt"
	"test/app/models"
)

func GetPlatFormData(query string) *[]models.NovelInfo {
	var novelInfoData *[]models.NovelInfo
	err := DB.Table("novel_info").Where("title LIKE ?", "%"+query+"%").Find(&novelInfoData)
	if err != nil {
		fmt.Println(err)
	}
	return novelInfoData
}

func GetPlatFormIntroduce(types string) *[]models.NovelInfo4 {
	var data *[]models.NovelInfo4
	var contentType string
	contentType = "none"
	if types == "navel" {
		contentType = "text"
	} else if types == "comic" {
		contentType = "png"
	}

	err := DB.Table("novel_info").Where("content_type = ?", contentType).Order("RAND()").Limit(6).Find(&data)
	if err != nil {
		fmt.Println(err)
	}

	return data
}
