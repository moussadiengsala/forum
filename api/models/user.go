package models

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Avatar    string `json:"avatar"`
	Password  string `json:"password"`
}
