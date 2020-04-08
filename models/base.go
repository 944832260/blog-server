package models

import "time"

type BaseStrcut struct {
	ID         int       `json:"id" gorm:"primary_key"`
	CreateTime time.Time `json:"create_time" sql:"default:now()" gorm:"type:DATETIME"`
}

type BaseModel struct {

}