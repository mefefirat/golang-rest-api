package users

import (
	"github.com/mefefirat/golang-rest-api/database"
)

type Users struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`

	/*CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`*/
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Add() {

}

func Edit() {

}

func Delete() {

}

func Get() {

}

func List() ([]Users, error) {

	rows, err := database.DB.Query("select * from users")
	defer rows.Close()
	checkError(err)
	var users []Users
	for rows.Next() {
		user := Users{}
		err := rows.Scan(&user.ID, &user.UserName)
		checkError(err)
		users = append(users, user)
	}

	return users, nil
}
