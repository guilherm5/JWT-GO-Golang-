package models

type Users struct {
	ID           int    `json:"id"`
	NameUser     string `json:"name_user"`
	EmailUser    string `json:"email_user"`
	PasswordUser string `json:"password_user"`
}

type InsertUsers struct {
	Records []Users `json:"records"`
}

type DeleteUsers struct {
	ID           []int  `json:"id"`
	NameUser     string `json:"name_user"`
	EmailUser    string `json:"email_user"`
	PasswordUser string `json:"password_user"`
}
