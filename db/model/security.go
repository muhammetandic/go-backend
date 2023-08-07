package model

type Role struct {
	Model
	Name       string    `json:"name" gorm:"unique;size:100" binding:"required"`
	Privileges Privilege `json:"privileges"`
	Users      []UserToRole
}

type UserToRole struct {
	Model
	UserID int `json:"userId" binding:"required"`
	RoleID int `json:"roleId" binding:"required"`
	Role   Role
	User   User
}

type Privilege struct {
	Model
	RoleID    int    `json:"roleId"`
	Endpoint  string `json:"endpoint" binding:"required"`
	CanRead   bool   `json:"canRead" gorm:"default:false"`
	CanInsert bool   `json:"canInsert" gorm:"default:false"`
	CanUpdate bool   `json:"canUpdate" gorm:"default:false"`
	CanDelete bool   `json:"canDelete" gorm:"default:false"`
}
