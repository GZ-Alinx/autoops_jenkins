package forms

type Jenkins_INFO struct {
	Url            string `form:"url" json:"url" binding:"required"`
	Login_user     string `form:"login_user" json:"login_user"`
	Login_password string `form:"login_password" json:"login_password"`
}

type Jenkins_form struct {
	Url            string `form:"url" json:"url" binding:"required"`
	Login_user     string `form:"login_user" json:"login_user"`
	Login_password string `form:"login_password" json:"login_password"`
}

type Jenkins_Get struct {
	Url string `form:"url" json:"url"`
}
