package entities

type User struct {
	UserId int `json:"userId" gorm:"primaryKey; autoincrement"`
	//PersonRefer int    `json:"personId" gorm:"primaryKey"`
	//RoleRefer   int    `json:"roleId" gorm:"primaryKey"`
	User     string `json:"user" gorm:"size:10; not null"`
	Password string `json:"password" gorm:"size:8; not null"`
	State    string `json:"state" gorm:"size:1; not null; default:'A'"`
	//Person      Person `gorm:"foreingKey:'PersonRefer'"`
	//Role        Role   `gorm:"foreingKey:'RoleRefer'"`
}
