package model

import (
	"time"

	"github.com/mumushuiding/util"
)

// CrewScheduling 员工排班表
type CrewSchedule struct {
	Model
	// 用户姓名
	Username string `json:"username"`
	// 公司
	Company string
	// 排班日期
	ScheduleDate string `json:"schedulingDate"`
	// 排班类型,如：早班
	Type string `json:"type"`
	// 排班人
	Creator string `json:"creator"`
	// 排班日期
	CreateTime string `json:"createTime"`
}

// Save Save
func (c *CrewSchedule) Save() error {
	c.CreateTime = util.FormatDate(time.Now(), util.YYYY_MM_DD_HH_MM_SS)
	return db.Create(c).Error
}

// DelCrewScheduleByID 根据ID删除
func DelCrewScheduleByID(id int) error {
	return db.Where("id=?", id).Delete(&CrewSchedule{}).Error
}
