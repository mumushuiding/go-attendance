package service

import (
	"github.com/mumushuiding/go-attendance/model"
)

// SaveCrewSchedule 保存用户的排班信息
func SaveCrewSchedule(c *model.CrewSchedule) error {
	return c.Save()
}

// DelCrewSchedule 根据id删除排班信息
func DelCrewSchedule(id int) error {
	return model.DelCrewScheduleByID(id)
}
