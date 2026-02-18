package model

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`   // Унікальний ідентифікатор
	Name        string    `gorm:"not null" json:"name"`   // Назва товару
	Description string    `json:"description"`            // Опис товару
	Price       float64   `gorm:"not null" json:"price"`  // Ціна
	Stock       int       `gorm:"default:0" json:"stock"` // Кількість на складі
	CreatedAt   time.Time `json:"created_at"`             // Дата створення
	UpdatedAt   time.Time `json:"updated_at"`             // Дата останнього оновлення
}
