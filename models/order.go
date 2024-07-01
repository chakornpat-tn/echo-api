package models

import "time"

type Order struct {
	ID          int32          `gorm:"primaryKey" json:"id"`
	UserID      int32          `json:"user_id"`
	AddressID   int32          `json:"address_id"`
	ShopID      int32          `json:"shop_id"`
	TotalPrice  float32        `gorm:"not null" json:"total_price"`
	CourierID   int32          `json:"courier_id"`
	OrderStatus int32          `json:"order_status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Products    []ProductOrder `gorm:"foreignKey:OrderID" json:"products"`
}

type OrderStatus struct {
	ID     int32   `gorm:"primaryKey" json:"id"`
	Title  string  `gorm:"not null" json:"title"`
	Orders []Order `gorm:"foreignKey:OrderStatus" json:"orders"`
}

type Courier struct {
	ID    int32   `gorm:"primaryKey" json:"id"`
	Title string  `gorm:"not null" json:"title"`
	Order []Order `gorm:"foreignKey:CourierID" json:"courier"`
}
