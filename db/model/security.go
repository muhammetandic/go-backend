package model

type Role struct {
	Model
	Name       string `json:"roleName" binding:"required"`
	Privileges []RolePrivilege
	Users      []UserToRole
}

type UserToRole struct {
	Model
	UserID int `json:"userId" binding:"required"`
	RoleID int `json:"roleId" binding:"required"`
	Role   Role
	User   User
}

type RolePrivilege struct {
	Model
	RoleID    int    `json:"roleId" binding:"required"`
	Table     string `json:"table" binding:"required"`
	CanRead   bool   `json:"canRead" binding:"required" gorm:"default:false"`
	CanInsert bool   `json:"canInsert" binding:"required" gorm:"default:false"`
	CanUpdate bool   `json:"canUpdate" binding:"required" gorm:"default:false"`
	CanDelete bool   `json:"CanDelete" binding:"required" gorm:"default:false"`
}
