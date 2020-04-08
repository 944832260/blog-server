package models

type User struct {
	BaseStrcut
	Tel     string `json:"tel" gorm:"size:11;not null;"`
	Name    string `json:"name" gorm:"size:32"`
	PassWord string `json:"password" gorm:"size:16"`

}

func (User) TableName() string  {
	return "users"
}

