package wechat_login

import (
	"encoding/json"
	"fmt"
	"github.com/BigBugsTech/bigbugs-wechat-utils/model"
	"github.com/BigBugsTech/bigbugs-wechat-utils/utils"
)

func GetAccessToken(appid, secret, code string) (res interface{}, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?"+
		"appid=%s&secret=%s&code=%s&grant_type=authorization_code", appid, secret, code)
	resBytes, errRes := utils.SendRequest(url, nil, nil, "get")
	if errRes != nil {
		err = errRes
		return nil, err
	}

	var accessToken model.AccessTokenResp
	var errMsg model.ErrMsgResp

	if errUnmarshal := json.Unmarshal(resBytes, &accessToken); errUnmarshal != nil { //if res is err,turn to unmarshal to err struct
		errUnmarshal = json.Unmarshal(resBytes, &errMsg)
		if errUnmarshal != nil { //if unmarshal indeed is err
			err = errUnmarshal
			return nil, errUnmarshal
		} else { // get errmsg
			res = errMsg
			err = nil
			return
		}
	} else { // get token successfully
		res = accessToken
		err = nil
		return
	}
}
