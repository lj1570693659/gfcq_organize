package entity

type WechatCryptError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WechatCryptToken struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

type HttpWechatDepart struct {
	ErrCode      int               `json:"errcode"`
	ErrMsg       string            `json:"errmsg"`
	DepartmentId []WechatDepartIds `json:"department_id"`
}

type WechatDepartIds struct {
	Id       int `json:"id"`
	ParentId int `json:"parentid"`
	Order    int `json:"order"`
}

type WechatDepart struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	NameEn           string   `json:"name_en"`
	DepartmentLeader []string `json:"department_leader"`
	ParentId         int      `json:"parentid"`
	Order            int      `json:"order"`
}

type HttpWechatDepartInfo struct {
	ErrCode    int              `json:"errcode"`
	ErrMsg     string           `json:"errmsg"`
	Department WechatDepartInfo `json:"department"`
}
type WechatDepartInfo struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	NameEn           string   `json:"name_en"`
	DepartmentLeader []string `json:"department_leader"`
	ParentId         int      `json:"parentid"`
	Order            int      `json:"order"`
}

// HttpWechatUser /*********获取部门成员详情***********/
type HttpWechatUser struct {
	ErrCode  int        `json:"errcode"`
	ErrMsg   string     `json:"errmsg"`
	UserList []UserInfo `json:"userlist"`
}
type UserInfo struct {
	UserId           string          `json:"userid"`
	Name             string          `json:"name"`
	Department       []int           `json:"department"`
	Order            []int           `json:"order"`
	Position         string          `json:"position"` // 职务信息
	Mobile           string          `json:"mobile"`
	Gender           string          `json:"gender"`
	Email            string          `json:"email"`
	BizMail          string          `json:"biz_mail"`      // 企业邮箱
	IsLeader         int             `json:"isleader"`      // 表示在所在的部门内是否为部门负责人。0-否；1-是
	DirectLeader     []string        `json:"direct_leader"` // 直属上级UserID
	Avatar           string          `json:"avatar"`
	Telephone        string          `json:"telephone"`
	Alias            string          `json:"alias"`
	Status           int             `json:"status"` // 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业。
	Address          string          `json:"address"`
	EnglishName      string          `json:"english_name"`
	OpenUserid       string          `json:"open_userid"`
	QrCode           string          `json:"qr_code"`           // 员工个人二维码
	ExternalPosition string          `json:"external_position"` // 对外职务
	Extattr          UserExtattrInfo `json:"extattr"`
}
type UserExtattrInfo struct {
	Attrs []UserAttrsInfo `json:"attrs"`
}
type UserAttrsInfo struct {
	Type  int    `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// HttpWechatCheckInDayDataReq /*********获取打卡日报***********/
type HttpWechatCheckInDayDataReq struct {
	StartTime  uint32   `json:"starttime"`
	EndTime    uint32   `json:"endtime"`
	UseridList []string `json:"useridlist"`
}

type HttpWechatCheckInDayDataRes struct {
	ErrCode int       `json:"errcode"`
	ErrMsg  string    `json:"errmsg"`
	Datas   []DayData `json:"datas"`
}

type DayData struct {
	SummaryInfo SummaryInfo `json:"summary_info"`
	OtInfo      OtInfo      `json:"ot_info"`
	BaseInfo    BaseInfo    `json:"base_info"`
}

// SummaryInfo 汇总信息
type SummaryInfo struct {
	CheckinCount    int32 `json:"checkin_count"`
	RegularWorkSec  int32 `json:"regular_work_sec"`
	StandardWorkSec int32 `json:"standard_work_sec"`
	EarliestTime    int32 `json:"earliest_time"`
	LastestTime     int32 `json:"lastest_time"`
	RecordType      int32 `json:"record_type"`
}

// OtInfo 加班信息
type OtInfo struct {
	OtStatus              int      `json:"ot_status"`
	OtDuration            int      `json:"ot_duration"`
	ExceptionDuration     []uint32 `json:"exception_duration"`
	WorkdayOverAsVacation int      `json:"workday_over_as_vacation"`
	WorkdayOverAsMoney    int      `json:"workday_over_as_money"`
}

type BaseInfo struct {
	Date        int      `json:"date"`
	RecordType  int      `json:"record_type"` // 记录类型：1-固定上下班；2-外出（此报表中不会出现外出打卡数据）；3-按班次上下班；4-自由签到；5-加班；7-无规则
	Name        string   `json:"name"`
	NameEx      string   `json:"name_ex"`
	DepartsName string   `json:"departs_name"`
	Acctid      string   `json:"acctid"`    // 打卡人员账号，即userid
	RuleInfo    RuleInfo `json:"rule_info"` // 打卡规则
	DayType     int      `json:"day_type"`  // 日报类型：0-工作日日报；1-休息日日报
}

// RuleInfo 打卡人员所属规则信息
type RuleInfo struct {
	Groupid      int    `json:"groupid"`
	Groupname    string `json:"groupname"`
	Scheduleid   int    `json:"scheduleid"`
	Schedulename string `json:"schedulename"`
	Checkintime  []struct {
		WorkSec    int `json:"work_sec"`     // 上班时间，为距离0点的时间差
		OffWorkSec int `json:"off_work_sec"` // 下班时间，为距离0点的时间差
	} `json:"checkintime"`
}

// SendImgMsgApiReq 发送图片消息请求参数
type SendImgMsgApiReq struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Image   struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
	Safe                   int `json:"safe"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// SendTextMsgApiReq 发送文本消息请求参数
type SendTextMsgApiReq struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe                   int `json:"safe"`
	EnableIdTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}
