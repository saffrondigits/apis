package modules

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}

type LoginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
