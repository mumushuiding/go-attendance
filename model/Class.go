package model

import (
	"time"

	"github.com/mumushuiding/util"
)

// Class 排班类型
type Class struct {
	Model
	Sections *Section
	// 为true,则下班不打卡也不会生成异常
	IsOffDutyFreeCheck bool `json:"isOffDutyFreeCheck"`
	// 班次名称，如："早班"
	Serial string
	// 旷工迟到分钟数
	AbsentTime int `json:"absentTime"`
	// 允许迟到分钟数
	PermitLateMinutes int `json:"permitLateMinutes"`
	// 严重迟到分钟数
	SeriousLateTime int `json:"seriousLateTime"`
	// 休息时间是否打开
	IsRestTimeOpen bool `json:"isRestTimeOpen"`
	// 工作时间，单位是分钟
	WorkTime int `json:"workTime"`
	// 休息时间
	RestTime Section `json:"restTime"`
	// 公司
	Company string
	// 创建人
	UserName string `json:"userName"`
	// 创建日期
	CreateTime string `json:"createTime"`
}

// ClassInfo 班次信息
type ClassInfo struct {
	// 班次名称，如："早班"
	Name string
	// 打卡时间段,如："10:00-16:30 "
	SectionsLabel string `json:"sectionsLabel"`
}

// Section 打卡时间段
type Section struct {
	Start CheckTime
	End   CheckTime
}

// CheckTime 打卡时间
type CheckTime struct {
	// 打卡时间，比如:07:00
	Time     string
	beginMin int
	endMin   int
}

// Save Save
func (s *Class) Save() error {
	s.CreateTime = util.FormatDate(time.Now(), util.YYYY_MM_DD_HH_MM_SS)
	return db.Create(s).Error
}

// Update Update
func (s *Class) Update() error {
	return db.Model(&Class{}).Updates(s).Error
}

// DelClassByID 根据ID删除
func DelClassByID(id int) error {
	return db.Where("id=?", id).Delete(&Class{}).Error
}
