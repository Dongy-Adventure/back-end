package model

type Seller struct {
	SellerID string `json:"sellerID" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password" copier:"-"`
	Name     string `json:"name" bson:"name"`
	Surname  string `json:"surname" bson:"surname"`
	Payment  string `json:"payment" bson:"payment"`
}
