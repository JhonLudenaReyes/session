package entities

type User struct {
	UserId   int    `json:"userId" gorm:"primaryKey; autoincrement"`
	PersonId int    `json:"personId" gorm:"primaryKey"`
	RoleId   int    `json:"RoleId" gorm:"primaryKey"`
	User     string `json:"user" gorm:"size:13; not null"`
	Password string `json:"password" gorm:"size:8; not null"`
	State    string `json:"state" gorm:"size:1; not null; default:'A'"`
	Person   Person `gorm:"foreingKey:'PersonId'"`
	Role     Role   `gorm:"foreingKey:'RoleId'"`
}
