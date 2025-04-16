package model

import (
	"time"
)

type Order struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      // 關聯到下訂的使用者
	Amount     float64   // 訂單總金額
	Status     string    // 訂單狀態，例如 "pending", "paid", "shipped"
	CreatedAt  time.Time
	UpdatedAt  time.Time
}