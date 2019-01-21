package fengcao

type AuthHeader struct {
	UserName    string
	Password    string
	Token       string
	Target      string
	AccessToken string
	Action      string // defalut API-SDK
}
