package types

type UserServiceReq struct {
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password"`
}
