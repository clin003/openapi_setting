package app_conf

var UserInfo UserInfoModel

type UserInfoModel struct {
	Username string `json:"mobile"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Action   string `json:"action"` //adduserinfo
}
