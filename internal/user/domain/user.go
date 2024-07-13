package domain

import (
	"time"
)

type Role string

const (
	RoleDeveloper        Role = "developer"
	RoleCustomer         Role = "customer"
	RoleDeveloperInvited Role = "developer_invited"
	RoleCustomerInvited  Role = "customer_invited"
)

type User struct {
	ID           int64      `json:"id" validate:"required,int"`
	PublicID     string     `json:"public_id" validate:"required,uuid"`
	Name         string     `json:"name" validate:"required,string,min=3,max=255"`
	Email        string     `json:"email" validate:"required,email"`
	Password     string     `json:"password" validate:"required,string,min=6,max=255"`
	LoginEnabled bool       `json:"login_enabled" validate:"required,bool"`
	Role         Role       `json:"role" validate:"required,role"`
	CreatedAt    time.Time  `json:"created_at" validate:"required,datetime"`
	UpdatedAt    time.Time  `json:"updated_at" validate:"required,datetime"`
	DeletedAt    *time.Time `json:"deleted_at" validate:"omitempty,datetime"`
}

func (u *User) Delete() {
	now := time.Now()
	u.DeletedAt = &now
}

func (u *User) IsDeleted() bool {
	return u.DeletedAt != nil
}

func (u *User) IsNotDeleted() bool {
	return u.DeletedAt == nil
}

func (u *User) EnableLogin() {
	u.LoginEnabled = true
}

func (u *User) DisableLogin() {
	u.LoginEnabled = false
}

func (u *User) IsLoginEnabled() bool {
	return u.LoginEnabled
}

func (u *User) IsNotLoginEnabled() bool {
	return !u.LoginEnabled
}

func (u *User) IsRole(role Role) bool {
	return u.Role == role
}
