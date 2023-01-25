package models


type User struct {
	ID	   int   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token	string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}



