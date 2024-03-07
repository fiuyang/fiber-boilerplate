package model

type Note struct {
	ID      int    `json:"id" gorm:"type:int;primary_key"`
	Content string `json:"content" gorm:"type:varchar(100)"`
}
