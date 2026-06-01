package model

type TASK struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Tasktx string `gorm:"not null" json:"tasktx"`
	Status bool `gorm:"default:false" json:"status"`
}