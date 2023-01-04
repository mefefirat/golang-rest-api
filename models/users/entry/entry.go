package entry

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}
