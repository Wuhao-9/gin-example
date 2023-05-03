package models

type Auth struct {
	ID      uint   `gorm:"primarykey;autoincrement" json:"id"`
	Account string `gorm:"type:varchar(20)" json:"account"`
	Passwd  string `gorm:"type:varchar(20)" json:"passwd"`
}
