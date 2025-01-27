package userrole

type UserType int

var UserRole = struct {
	BUYER  UserType
	SELLER UserType
}{
	BUYER:  0,
	SELLER: 1,
}
