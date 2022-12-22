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

func Create(username string) (int, error) {

	query := `INSERT INTO users ("username") VALUES ($1) RETURNING id`
	stmt, err := database.DB.Prepare(query)
	if stmt != nil {
		defer stmt.Close()
	}

	var id int
	err = stmt.QueryRow(username).Scan(&id)
	if err != nil {
		panic(err)
	}

	return id, nil

	/*sql := "INSERT INTO users (username) VALUES ($1);"
	result, err := database.DB.Exec(sql, "test")*/

}

func List() (*[]Users, error) {

	rows, err := database.DB.Query("select * from users")
	if rows != nil {
		defer rows.Close()
	}

	users := []Users{}

	if err != nil {
		return &[]Users{}, err
	}
	for rows.Next() {
		user := Users{}
		err := rows.Scan(&user.ID, &user.UserName)
		if err != nil {
			return &[]Users{}, err
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
