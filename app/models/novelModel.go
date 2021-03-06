package models

import "time"

type Novel struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Page        int    `json:"page"`
	NovelNameId int    `json:"novel_name_id"`
}

type NovelContent struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Page        int    `json:"page"`
	NovelNameId int    `json:"novel_name_id"`
}

type NovelList struct {
	ID          int    `json:"id"`
	NovelName   string `json:"novel_name"`
	NovelNameId int    `json:"novel_name_id"`
}

type NovelInfo struct {
	ID             int       `gorm:"AUTO_INCREMENT"`
	Title          string    `json:"title"`
	Author         string    `json:"author"`
	NovelNameId    string    `json:"novel_name_id"`
	ComicNameId    string    `json:"comic_name_id"`
	LongInfo       string    `json:"long_info"`
	Tags           string    `json:"tags"`
	Cat            string    `json:"cat"`
	ContentType    string    `json:"content_type"`
	TitlePhotoUrl  string    `json:"title_photo_url"`
	DataUpdateTime time.Time `json:"data_update_time"`
}

type NovelInfo2 struct {
	ID          int    `gorm:"AUTO_INCREMENT"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	NovelNameId int    `json:"novel_name_id"`
	ComicNameId int    `json:"comic_name_id"`
	LongInfo    string `json:"long_info"`
	Tags        string `json:"tags"`
	Cat         string `json:"cat"`
	ContentType string `json:"content_type"`
	// TitlePhotoUrl  string    `json:"title_photo_url"`
	DataUpdateTime time.Time `json:"data_update_time"`
	Novel          Novel     `gorm:"foreignKey:NovelNameId;references:NovelNameId;"`
}

func (Novel) TableName() string { return "novel" }

type NovelInfo4 struct {
	ID             int       `gorm:"AUTO_INCREMENT"`
	Title          string    `json:"title"`
	Author         string    `json:"author"`
	NovelNameId    int       `json:"novel_name_id"`
	ComicNameId    string    `json:"comic_name_id"`
	LongInfo       string    `json:"long_info"`
	Tags           string    `json:"tags"`
	Cat            string    `json:"cat"`
	ContentType    string    `json:"content_type"`
	TitlePhotoUrl  string    `json:"title_photo_url"`
	DataUpdateTime time.Time `json:"data_update_time"`
	Page           int       `json:"page"`
	PageTitle      string    `json:"page_title"`
}

type BookList struct {
	ID          int    `gorm:"AUTO_INCREMENT"`
	Title       string `json:"title"`
	Page        int    `json:"page"`
	Name        int    `json:"name"`
	NovelNameId int    `json:"novel_name_id"`
}
