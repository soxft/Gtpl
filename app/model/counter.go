package model

type Counter struct {
	Count    int   `gorm:"type:int;not null;default:0"`
	UpdateAt int64 `gorm:"autoUpdateTime"`
}
