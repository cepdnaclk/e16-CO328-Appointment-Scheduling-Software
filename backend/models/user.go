package models

type SignupUser struct{
	FristName string  `json:"fristName"`
	LastName string  `json:"lastName"`
	Email string  `json:"email"`
	WhatsAppNo string `json:"phoneNo"`
	Password string `json:"password"`
};

type LoginUser struct{
	Email string `json:"email"`
	Password string `json:"password"`
}