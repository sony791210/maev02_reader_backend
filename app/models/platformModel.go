package models

import "time"

type PlatformInfo struct {
	ID             int       `gorm:"AUTO_INCREMENT"`
	Title          string    `json:"title"`
	Author         string    `json:"author"`
	NovelNameId    string    `json:"novel_name_id"`
	ComicNameId    string    `json:"comic_name_id"`
	VideoNameId    string    `json:"video_name_id"`
	LongInfo       string    `json:"long_info"`
	Tags           string    `json:"tags"`
	Cat            string    `json:"cat"`
	ContentType    string    `json:"content_type"`
	TitlePhotoUrl  string    `json:"title_photo_url"`
	DataUpdateTime time.Time `json:"data_update_time"`
}
