package models

type RefreshToken struct {
	User_id     uint
	Token       string
	Expiry_date string
	Active      bool
}
