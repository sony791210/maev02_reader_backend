package models

import "time"

type ComicInfo4 struct {
	ID             int       `gorm:"AUTO_INCREMENT"`
	Title          string    `json:"title"`
	Author         string    `json:"author"`
	ComicNameId    string    `json:"comic_name_id"`
	LongInfo       string    `json:"long_info"`
	Tags           string    `json:"tags"`
	Cat            string    `json:"cat"`
	ContentType    string    `json:"content_type"`
	TitlePhotoUrl  string    `json:"title_photo_url"`
	DataUpdateTime time.Time `json:"data_update_time"`
}

type ComicListInfo struct {
	FilePath []string `json:"filePath"`
	Title    string   `json:"title"`
	Page     int      `json:"page"`
}

type Comic struct {
	ID    int    `gorm:"AUTO_INCREMENT"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Page  string `json:"page"`
}
