package entities

type Role struct {
	RoleId int    `json:"roleId" gorm:"primaryKey; autoincrement"`
	Role   string `json:"role" gorm:"size:15; not null"`
	State  string `json:"state" gorm:"size:1; not null; default:'A'"`
}
