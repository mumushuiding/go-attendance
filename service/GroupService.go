package service

import (
	"github.com/mumushuiding/go-attendance/model"
)

// GroupReceiver GroupReceiver
type GroupReceiver struct {
	// 考核组名称
	Name string
	// 类型有：FIXED(固定班制),TURN(轮班制)
	Type string
	// 打卡有效范围,单位米
	Offset int
	// 默认班次
	DefaultClassID int `json:"defaultClassId"`
	// 考核组用户
	Users []string
	// 管理员
	Managers []string
	// 无需打卡人员
	ExcludeUsers []string `json:"excludeUsers"`
	// 班次
	ClassesIds []int `json:"classesIds"`
	// 打卡地点
	Locations []*model.Location
	// 根据WiFi考勤
	Wifis []*model.Wifi
	// 工作日,如 [0,362510019,362510019,362510019,362510019,362510019,0]，0表示不上班，362510019为班次Id
	WorkDayList  []int `json:"workDayList"`
	WorkDays     []int `json:"workDays"`
	SkipHolidays bool  `json:"skipHolidays"`
}
