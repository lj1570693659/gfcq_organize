package user

import (
	"errors"
	"fmt"
	"github.com/gogf/gf-demos/app/model/user"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gvalid"
)

const (
	USER_SESSION_MARK = "user_info"
)

// 用户注册
func SignUp(data g.MapStrStr) error {
	// 数据校验
	rules := []string{
		"passport @required|length:6,16#账号不能为空|账号长度应当在:min到:max之间",
		"password2@required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间",
		"password @required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等",
	}
	if e := gvalid.CheckMap(data, rules); e != nil {
		return errors.New(e.String())
	}
	if _, ok := data["nickname"]; !ok {
		data["nickname"] = data["passport"]
	}
	// 唯一性数据检查
	if !CheckPassport(data["passport"]) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", data["passport"]))
	}
	if !CheckNickName(data["nickname"]) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", data["nickname"]))
	}
	// 记录账号创建/注册时间
	if _, ok := data["create_time"]; !ok {
		data["create_time"] = gtime.Now().String()
	}
	if _, err := user.Model.Filter().Data(data).Save(); err != nil {
		return err
	}
	return nil
}

// 判断用户是否已经登录
func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains(USER_SESSION_MARK)
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func SignIn(passport, password string, session *ghttp.Session) error {
	one, err := user.FindOne("passport=? and password=?", passport, password)
	if err != nil {
		return err
	}
	if one == nil {
		return errors.New("账号或密码错误")
	}
	session.Set(USER_SESSION_MARK, one)
	return nil
}

// 用户注销
func SignOut(session *ghttp.Session) {
	session.Remove(USER_SESSION_MARK)
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckPassport(passport string) bool {
	if i, err := user.FindCount("passport", passport); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckNickName(nickname string) bool {
	if i, err := user.FindCount("nickname", nickname); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 获得用户信息详情
func GetProfile(session *ghttp.Session) (u *user.Entity) {
	session.GetStruct(USER_SESSION_MARK, &u)
	return
}
