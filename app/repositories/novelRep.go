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

func GetInfoData(id int) interface{} {
	var novelInfoData models.NovelInfo4

	err := DB.Table("novel_info").Select("novel_info.*,novel.page,novel.title as page_title").
		Joins("left join novel on novel.novel_name_id = novel_info.novel_name_id").
		Where("novel.page = (SELECT MAX(page) from novel where novel_name_id =?) ", id).
		Where("novel_info.novel_name_id = ?", id).First(&novelInfoData)

	if err != nil {
		fmt.Println(err)
	}

	return novelInfoData
}

func GetDetailData(id int, page int) interface{} {
	var novelDetailData models.NovelContent
	err := DB.Table("novel").Where("novel_name_id =?", id).Where("page = ?", page).First(&novelDetailData)
	if err != nil {
		fmt.Println(err)
	}
	return novelDetailData
}
