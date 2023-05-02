package models

// define model for Article
type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag" gorm:"foreignkey:TagID"`

	Title      string `json:"title" gorm:"type:varchar(50)"`
	Desc       string `json:"desc" gorm:"type:text"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by" gorm:"type:varchar(20)"`
	ModifiedBy string `json:"modified_by" gorm:"type:varchar(20)"`
	State      int    `json:"state" gorm:"type:tinyint UNSIGNED"`
}
