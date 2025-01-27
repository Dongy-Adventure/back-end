package model

type Buyer struct {
	BuyerID string `json:"buyerID" bson:"buyer_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password" copier:"-"`
	Name     string `json:"name" bson:"name"`
	Surname  string `json:"surname" bson:"surname"`
}
