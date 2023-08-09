package models

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string
	Email    string
}
