package users

import (
	"github.com/mefefirat/golang-rest-api/database"
	"github.com/mefefirat/golang-rest-api/models/users/entry"
)

func Create(user *entry.User) (int, error) {

	query := `INSERT INTO users("username", "email", "first_name", "last_name") VALUES ($1, $2, $3, $4) RETURNING id`
	stmt, err := database.DB.Prepare(query)
	if stmt != nil {
		defer stmt.Close()
	}

	var id int
	err = stmt.QueryRow(user.UserName, user.Email, user.FirstName, user.LastName).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil

	/*sql := "INSERT INTO users (username) VALUES ($1);"
	result, err := database.DB.Exec(sql, "test")*/

}

func List() (*[]entry.User, error) {

	rows, err := database.DB.Query("select * from users")
	if rows != nil {
		defer rows.Close()
	}

	users := []entry.User{}

	if err != nil {
		return &[]entry.User{}, err
	}
	for rows.Next() {
		user := entry.User{}
		err := rows.Scan(&user.ID, &user.UserName, &user.Email)
		if err != nil {
			return &[]entry.User{}, err
		}
		users = append(users, user)
	}

	return &users, nil
}

func Edit() {

}

func Delete() {

}

func Get() {

}
