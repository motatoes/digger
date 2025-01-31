// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID             string         `gorm:"column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Email          string         `gorm:"column:email;not null" json:"email"`
	OrganisationID string         `gorm:"column:organisation_id" json:"organisation_id"`
	ExternalID     string         `gorm:"column:external_id" json:"external_id"`
	FirstName      string         `gorm:"column:first_name" json:"first_name"`
	LastName       string         `gorm:"column:last_name" json:"last_name"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
