package entities

import "time"

type Admin struct {
	ID       uint      `gorm:"primary_key"`
	Username string    `gorm:"column:username"`
	Password string    `gorm:"column:password"`
	RoleID   int       `gorm:"column:role_id"`
	Verified string    `gorm:"column:verified"`
	Active   string    `gorm:"column:active"`
	CreateAt time.Time `gorm:"column:created_at"`
	UpdateAt time.Time `gorm:"column:updated_at"`
}

type RegistrationApproval struct {
	ID         uint      `gorm:"primary_key"`
	AdminID    uint      `gorm:"column:admin_id"`
	UserID     uint      `gorm:"column:user_id"`
	Status     string    `gorm:"column:status"` // pending, approved, rejected
	ApprovedAt time.Time `gorm:"column:approved_at"`
	RejectedAt time.Time `gorm:"column:rejected_at"`
}

func (Admin) TableName() string {
	return "actors"
}

type Role struct {
	ID       int    `gorm:"primary_key"`
	RoleName string `gorm:"column:role_name"`
}
