package repositories

import (
	"fmt"
	"test/app/models"
)

func GetPlatFormData(query string) *[]models.PlatformInfo {
	var PlatformInfoData *[]models.PlatformInfo
	err := DB.Table("platform_info").Where("title LIKE ?", "%"+query+"%").Find(&PlatformInfoData)
	if err != nil {
		fmt.Println(err)
	}
	return PlatformInfoData
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

	err := DB.Table("novel_info").Where("content_type = ?", contentType).Order("RAND()").Limit(3).Find(&data)
	if err != nil {
		fmt.Println(err)
	}

	return data
}

func GetPlatFormLastInfo() *[]models.NovelInfo4 {
	var data *[]models.NovelInfo4

	err := DB.Table("novel_info").Order("update_time DESC").Limit(4).Find(&data)
	if err != nil {
		fmt.Println(err)
	}

	return data
}

func GetPlatFormList() interface{} {
	var PlatformInfoPngData []models.PlatformInfo
	var PlatformInfoTxtData []models.PlatformInfo
	var PlatformInfoMp4Data []models.PlatformInfo

	err := DB.Table("platform_info").Where("content_type='png'").Order("update_time DESC").Limit(5).Find(&PlatformInfoPngData)
	if err != nil {
		fmt.Println(err)
	}
	err = DB.Table("platform_info").Where("content_type='txt'").Order("update_time DESC").Limit(5).Find(&PlatformInfoTxtData)
	if err != nil {
		fmt.Println(err)
	}
	err = DB.Table("platform_info").Where("content_type='mp4'").Order("update_time DESC").Limit(5).Find(&PlatformInfoMp4Data)
	if err != nil {
		fmt.Println(err)
	}

	PlatformInfoAllData := append(PlatformInfoPngData, PlatformInfoTxtData...)
	PlatformInfoAllData = append(PlatformInfoAllData, PlatformInfoMp4Data...)

	return PlatformInfoAllData
}
