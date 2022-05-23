package model

type AccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid string `json:"openid"`
	Scope string `json:"scope"`
	Unionid string `json:"unionid"`
}

type RefreshTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid string `json:"openid"`
	Scope string `json:"scope"`
}

type UserInfoResp struct{
	Openid string `json:"openid"`
	Nickname string `json:"nickname"`
	Sex int64 `json:"sex"`
	Province string `json:"province"`
	City string `json:"city"`
	Country string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	Privilege string `json:"privilege"`
	Unionid string  `json:"unionid"`
}

type ErrMsgResp struct {
	Errcode string `json:"errcode"`
	Errmsg string `json:"errmsg"`
}



