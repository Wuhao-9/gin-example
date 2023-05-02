package models

// define model for Tag
type Tag struct {
	Model
	Name       string `json:"name" gorm:"type:varchar(20);unique"`
	CreatedBy  string `json:"created_by" gorm:"type:varchar(20)"`
	ModifiedBy string `json:"modified_by" gorm:"type:varchar(20)"`
	State      int    `json:"state" gorm:"type:tinyint UNSIGNED"`
}
