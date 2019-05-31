package model

// AttendanceDetail 用户考勤明细
type AttendanceDetail struct {
	ApproveList    []string   `json:"approveList"`
	AttendDayTimes float64    `json:"attendDayTimes"`
	ClassInfo      *ClassInfo `json:"classInfo"`
	UserName       string     `json:"userName"`
	Compapny       string     `json:"company"`
	// 考核组
	Group *Group
	// 考核组名称
	GroupName string `json:"groupName"`
	// 是否有查阅权限
	HasViewPermission bool `json:"hasViewPermission"`
	// 是否能够编辑
	IsEditable bool `json:"isEditable"`
	// 是否有排班
	IsHasSchedule bool `json:"isHasSchedule"`
	// 是否休息
	IsRest bool `json:"isRest"`
	// 打卡纪录
	RecordList []*Record `json:"recordList"`
	// 工作时常，单位分钟
	WorkTime int `json:"workTime"`
}
