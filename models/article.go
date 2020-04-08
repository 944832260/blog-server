package models

type Article struct {
	BaseStrcut
	Title		string	`json:"title" gorm:"type:varchar(200);"`
	Type		string	`json:"type" `
	Status		int		`json:"status" `
	Content 	string	`json:"content" gorm:"type:longtext;"`
	AuthorId 	int		`json:"author_id" `
}

func (Article) TableName() string  {
	return "articles"
}

