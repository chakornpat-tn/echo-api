package models

import "time"

type Product struct {
	ID         int32                `gorm:"primaryKey" json:"id"`
	ShopID     int32                `json:"shop_id"`
	Title      string               `gorm:"not null" json:"title"`
	Price      float32              `gorm:"not null" json:"price"`
	Detail     string               `gorm:"not null" json:"detail"`
	CreatedAt  time.Time            `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time            `gorm:"not null" json:"updated_at"`
	Categories []ProductsCategories `gorm:"foreignKey:ProductID" json:"categories"`
	Orders     []ProductOrder       `gorm:"foreignKey:ProductID" json:"orders"`
}

type Category struct {
	ID       int32                `gorm:"primaryKey" json:"id"`
	Title    string               `json:"title"`
	Products []ProductsCategories `gorm:"foreignKey:CategoryID" json:"products"`
}

type ProductsCategories struct {
	ProductID  int32 `json:"product_id"`
	CategoryID int32 `json:"category_id"`
}

type ProductOrder struct {
	ProductID int32 `json:"product_id"`
	OrderID   int32 `json:"order_id"`
	Qty       int16 `gorm:"not null" json:"qty"`
}
