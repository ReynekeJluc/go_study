package models

import "time"

type Token struct {
	TokenId   uint      `gorm:"primaryKey;column:token_id;autoIncrement"`
	UserId    uint      `gorm:"column:user_id"`
	User      User      `gorm:"foreignKey:UserId;references:UserId;constraint:OnDelete:CASCADE"`
	TokenHash string    `gorm:"column:token_hash;size:64;uniqueIndex"`
	Revoked   bool      `gorm:"column:revoked;default:false"`
	ExpiresAt time.Time `gorm:"column:expires_at"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Token) TableName() string {
	return "tokens"
}
