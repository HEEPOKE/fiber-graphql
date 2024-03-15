package jwt

type JwtPayloadModel struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Exp      int64  `json:"exp"`
}
