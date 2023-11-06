package dao

var (
	GetTokenUrl                = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	GetDepartmentListUrl       = "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s&id=%d"
	GetDepartmentSimpleListUrl = "https://qyapi.weixin.qq.com/cgi-bin/department/simplelist?access_token=%s&id=%d"
	// GetDepartmentInfoUrl 获取单个部门详情
	GetDepartmentInfoUrl = "https://qyapi.weixin.qq.com/cgi-bin/department/get?access_token=%s&id=%d"

	// GetUserLists /***成员信息API****/
	GetUserLists        = "https://qyapi.weixin.qq.com/cgi-bin/user/list_id?access_token=%s"
	GetUserListByDepart = "https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=%s&department_id=%d"

	// GetUserCheckInDayData /***成员打卡API****/
	GetUserCheckInDayData = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckin_daydata?access_token=%s"

	// SendMsgApi /***发送应用消息****/
	SendMsgApi = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
)
