package main

type UserCrendetial struct {
	Username string		`json:"user_name"` // `form:"username" binding:"required"`
	Password string		`json:"user_password"` //form:"password" binding:"required"`
}

type AuthHeader struct {
	AuthorizationHeader		string		`header:"Authorization"`
}