package model

type Product struct {
	Id            int    `gorm:"type:int;primary_key"`
	Name          string `gorm:"type:varchar(255)"`
	Category      string `gorm:"type:varchar(255)"`
	OriginalPrice string `gorm:"type:varchar(255)"`
	DiscountPrice string `gorm:"type:varchar(255)"`
	Supermarket   string `gorm:"type:varchar(255)"`
}
