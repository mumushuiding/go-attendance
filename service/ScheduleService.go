package service

import (
	"github.com/mumushuiding/go-attendance/model"
)

// SaveSchedule 保存用户的排班信息
func SaveSchedule(c *model.Schedule) error {
	return c.Save()
}

// DelSchedule 根据id删除排班信息
func DelSchedule(id int) error {
	return model.DelScheduleByID(id)
}
