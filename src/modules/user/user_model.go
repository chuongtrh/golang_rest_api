package user

import (
	"time"
)

const (
	RoleAdmin   string = "admin"
	RoleManager string = "manager"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Email     string    `gorm:"type:varchar(256);unique;unique_index;not null" json:"email"`
	Password  string    `gorm:"type:varchar(512)" json:"password"`
	Role      string    `gorm:"type:varchar(256)" json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (user *User) ToObject() map[string]interface{} {
	return map[string]interface{}{
		"id":        user.ID,
		"email":     user.Email,
		"role":      user.Role,
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
	}
}
