package models

import "time"

// type User struct {
// 	ID        uint      `gorm:"primaryKey" json:"id"`
// 	Name      string    `json:"name"`
// 	Email     string    `json:"email"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

type User struct {
	ID            int32         `gorm:"primaryKey" json:"id"`
	Name          string        `gorm:"not null" json:"name"`
	Email         string        `gorm:"not null" json:"email"`
	Password      string        `gorm:"not null" json:"password"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	UserAddresses []UserAddress `gorm:"foreignKey:UserID" json:"user_addresses"`
	Orders        []Order       `gorm:"foreignKey:UserID" json:"orders"`
}

type UserAddress struct {
	ID            int32   `gorm:"primaryKey"  json:"id"`
	UserID        int32   `json:"user_id"`
	RecipientName string  `gorm:"not null" json:"recipient_name"`
	TelNumber     string  `gorm:"not null" json:"tel_number"`
	ZipCode       string  `gorm:"not null" json:"zip_code"`
	Province      string  `gorm:"not null" json:"province"`
	AddressDetail string  `gorm:"not null" json:"address_detail"`
	Orders        []Order `gorm:"foreignKey:AddressID" json:"orders"`
}
