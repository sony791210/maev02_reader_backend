package repositories

import (
	"fmt"
	"test/app/models"
)

func GetNovelList() *[]models.NovelList {
	var result *[]models.NovelList
	err := DB.Table("novel_list").Find(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func GetNovelTitleList(novel_name_id int) *[]models.Novel {
	var result *[]models.Novel
	err := DB.Table("novel").Where("novel_name_id = ?", novel_name_id).Find(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func GetNovelContent(id int) models.NovelContent {
	var result models.NovelContent
	err := DB.Table("novel").Where("id = ?", id).Find(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func GetNovelByPage(page int) models.Novel {
	var novelData models.Novel
	err := DB.Where("page = ?", page).First(&novelData)
	if err != nil {
		fmt.Println(err)
	}
	return novelData
}

func GetAllData(query string) *[]models.NovelInfo {
	var novelInfoData *[]models.NovelInfo
	err := DB.Table("novel_info").Where("title LIKE ?", "%"+query+"%").Find(&novelInfoData)
	if err != nil {
		fmt.Println(err)
	}
	return novelInfoData
}

func GetInfoData(id int) models.NovelInfo {
	var novelInfoData models.NovelInfo
	err := DB.Table("novel_info").Where("novel_name_id = ?", id).First(&novelInfoData)
	if err != nil {
		fmt.Println(err)
	}
	return novelInfoData
}
