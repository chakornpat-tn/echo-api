package models

type Shop struct {
	ID        int32     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Address   string    `gorm:"not null" json:"address"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Products  []Product `gorm:"foreignKey:ShopID" json:"products"`
	Orders    []Order   `gorm:"foreignKey:ShopID" json:"orders"`
}
