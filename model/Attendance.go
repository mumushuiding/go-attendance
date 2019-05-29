package model

// Attendance 打卡纪录（统计异常纪录表)
type Attendance struct {
	Model
	// 用户姓名
	Username string `json:"username"`
	// 打卡时间，比如："2019/4/1 星期一 上午 11:06"
	ClockTime string `json:"clockTime"`
	// 打卡状态，比如：上班签到，下班签退
	ClockState string `json:"clockState"`
	// 更正状态
	CorrectState string `json:"correctState"`
	// 异常情况
	ExceptionState string `json:"exceptionState"`
	// 公司
	Company string `json:"company"`
}

// Save save
func (a *Attendance) Save() error {
	return db.Create(a).Error
}
