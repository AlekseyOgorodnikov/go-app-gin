package entity

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `gorm:"type:varchar(32)" json:"firstname" binding:"required"`
	LastName  string `gorm:"type:varchar(32)" json:"lastname" binding:"required"`
	Age       int    `gorm:"type:varchar(10)" json:"age" binding:"gte=1,lte=130"`
	Email     string `gorm:"type:varchar(256);UNIQUE" json:"email" binding:"required,email"`
}

type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"type:varchar(100)" json:"title" binding:"min=2,max=100" validate:"is-cool"`
	Description string    `gorm:"type:varchar(200)" json:"description" binding:"max=200"`
	URL         string    `gorm:"type:varchar(256);UNIQUE" json:"url"  binding:"required,url"`
	Author      Person    `gorm:"foreignkey:PersonID" json:"author"  binding:"required"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	PersonID    uint64    `json:"-"`
}
