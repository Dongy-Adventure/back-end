package tokenmode

type TokenType int

var TokenMode = struct {
	ACCESS_TOKEN  TokenType
	REFRESH_TOKEN TokenType
}{
	ACCESS_TOKEN:  0,
	REFRESH_TOKEN: 1,
}
