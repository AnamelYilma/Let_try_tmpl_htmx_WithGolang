package Model

type TASK struct {
	ID     uint `gorm:"PrimaryKey"`
	Tasktx string `gorm:"not null"`
	status bool `gorm:"default:false"`
}