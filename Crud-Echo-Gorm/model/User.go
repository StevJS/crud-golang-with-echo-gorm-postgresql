package model

type User struct {
	ID       int    `gorm:"primaryKey;auto_increment" json:"id"`
	Name     string `json:name`
	LastName string `json:last_name`
	Email    string `json:email`
	Age      int    `json:age`
}
