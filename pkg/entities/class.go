package entities

type Class struct {
	ClassName string `gorm:"size:120;not null"`
}
