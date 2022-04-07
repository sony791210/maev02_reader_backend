package models

import "time"


type ComicInfo4 struct {
	ID             int       `gorm:"AUTO_INCREMENT"`
	Title          string    `json:"title"`
	Author         string    `json:"author"`
	ComicNameId    string     `json:"comic_name_id"`
	LongInfo       string    `json:"long_info"`
	Tags           string    `json:"tags"`
	Cat            string    `json:"cat"`
	ContentType    string    `json:"content_type"`
	TitlePhotoUrl  string    `json:"title_photo_url"`
	DataUpdateTime time.Time `json:"data_update_time"`
}


type ComicListInfo struct {
	FilePath []string
	Title string
	page int
}