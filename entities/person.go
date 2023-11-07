package entities

type Person struct {
	PersonId  int    `json:"personId" gorm:"primaryKey; autoincrement"`
	Name      string `json:"name" gorm:"size:30; not null"`
	LastName  string `json:"lastName" gorm:"size:30; not null"`
	Email     string `json:"email" gorm:"size:80; not null"`
	CellPhone string `json:"cellPhone" gorm:"size:10; not null"`
	State     string `json:"state" gorm:"size:1; default:'A'"`
}
