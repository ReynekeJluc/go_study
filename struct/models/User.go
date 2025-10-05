package models

type User struct {
	UserId    uint   `gorm:"column:user_id;primaryKey;autoIncrement"`
	Login     string `gorm:"column:login;size:255;not null"`
	Password  string `gorm:"column:password;size:255;not null"`
	CreatedAt string `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}
