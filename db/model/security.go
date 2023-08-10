package model

type Role struct {
	Model
	Name       string       `json:"name" gorm:"unique;size:100" binding:"required"`
	Privileges []Privilege  `json:"privileges"`
	Users      []UserToRole `json:"users" binding:"-"`
}

type UserToRole struct {
	Model
	UserID int  `json:"userId" binding:"required"`
	RoleID int  `json:"roleId" binding:"required"`
	Role   Role `json:"role" binding:"-"`
	User   User `json:"user" binding:"-"`
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

type UserToRoleDto struct {
	ID     int `json:"id"`
	UserID int `json:"userId"`
	RoleID int `json:"roleId"`
}
