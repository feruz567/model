package model

type User struct {
	Id int           `json: "id"`
	Name string      `json: "name"`
	Username string  `json: "username"`
	Phone int        `json: "phone"`
}