package entities

type Person struct {
	PersonId int    `json:"personId" gorm:"primaryKey; autoincrement"`
	Name     string `json:"name" gorm:"size:30; not null"`
	LastName string `json:"lastName" gorm:"size:30; not null"`
	Addresss string `json:"address" gorm:"size:120; not null"`
	State    string `json:"state" gorm:"size:1; default:'A'"`
}
