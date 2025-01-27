package dto

type Buyer struct {
	BuyerID string `json:"buyerID"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}