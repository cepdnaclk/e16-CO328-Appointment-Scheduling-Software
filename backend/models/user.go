package models

type SignupUser struct{
	FristName string  `json:"fristName"`
	LastName string  `json:"lastLame"`
	Email string  `json:"email"`
	WhatsAppNo string `json:"whatsappNo"`
	Password string `json:"password"`
};

type LoginUser struct{
	Email string `json:"email"`
	Password string `json:"password"`
}