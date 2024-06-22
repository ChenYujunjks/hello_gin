package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type ArithmeticRequest struct {
	Number1 float64 `json:"number1" binding:"required"`
	Number2 float64 `json:"number2" binding:"required"`
}
