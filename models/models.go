package models

import (
	"time"

	"gorm.io/gorm"
)

type Hewan struct {
	ID           uint `gorm:"primaryKey"`
	Nama         string
	JenisKelamin string
	JenisHewan   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Produk struct {
	ID        uint `gorm:"primaryKey"`
	Nama      string
	Harga     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
